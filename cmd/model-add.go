package cmd

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/peterszarvas94/goat/config"
)

func modelAdd(modelname string) error {
	// 0. check -> TODO imporve
	if modelname == "" {
		return fmt.Errorf("Name can not be empty")
	}

	// 1. MIGRATIONS

	err := existOrCreateDir(config.MigrationsPath)
	if err != nil {
		return err
	}

	output, err := cmd("goose", "-dir", config.MigrationsPath, "create", fmt.Sprintf("create_%s_table", modelname), "sql")
	fmt.Println(output)
	if err != nil {
		return err
	}

	migrationFilepath, err := getFileNameFromGooseOutput(output)
	if err != nil {
		return err
	}

	migrationSQL := fmt.Sprintf(`-- +goose Up
CREATE TABLE %s (
  id TEXT PRIMARY KEY,
  name TEXT NOT NULL
);

-- +goose Down
DROP TABLE %s;`, modelname, modelname)

	err = os.WriteFile(migrationFilepath, []byte(migrationSQL), 0644)
	if err != nil {
		return err
	}

	fmt.Printf("Migration file created: %s\n", migrationFilepath)

	// 2. SCHEMA

	err = existOrCreateDir(config.SchemaDirPath)
	if err != nil {
		return err
	}

	schemaFilePath := filepath.Join(config.SchemaDirPath, fmt.Sprintf("%s.sql", modelname))
	err = existsOrCreateFile(schemaFilePath)
	if err != nil {
		return err
	}

	modelSQL := fmt.Sprintf(`CREATE TABLE %s (
  id TEXT PRIMARY KEY,
  name TEXT NOT NULL
);`,
		modelname,
	)

	err = os.WriteFile(schemaFilePath, []byte(modelSQL), 0644)
	if err != nil {
		return err
	}

	fmt.Printf("Schema is created: %s\n", schemaFilePath)

	// 3. QUERIES

	err = existOrCreateDir(config.QueriesDirPath)
	if err != nil {
		return err
	}

	queriesFilePath := filepath.Join(config.QueriesDirPath, fmt.Sprintf("%s.sql", modelname))
	err = existsOrCreateFile(queriesFilePath)
	if err != nil {
		return err
	}

	uppercasemodelName := fmt.Sprintf("%s%s", strings.ToUpper(string(modelname[0])), modelname[1:])

	queriesSQL := fmt.Sprintf(`-- name: Get%sByID :one
SELECT *
FROM %s
WHERE id = ?;

-- name: List%ss :many
SELECT *
FROM %s
ORDER BY name;

-- name: Create%s :one
INSERT INTO %s (id, name)
VALUES (?, ?)
RETURNING *;

-- name: Update%s :one
UPDATE %s
SET name = ?
WHERE id = ?
RETURNING *;

-- name: Delete%s :exec
DELETE FROM %s
WHERE id = ?;`,
		uppercasemodelName,
		modelname,
		uppercasemodelName,
		modelname,
		uppercasemodelName,
		modelname,
		uppercasemodelName,
		modelname,
		uppercasemodelName,
		modelname,
	)

	err = os.WriteFile(queriesFilePath, []byte(queriesSQL), 0644)
	if err != nil {
		return err
	}

	fmt.Printf("Queries are created: %s\n", queriesFilePath)

	output, err = cmd("sqlc", "generate")
	fmt.Println(output)
	if err != nil {
		return err
	}

	fmt.Println("Run migration with \"goat migrate:up\"")

	return nil
}

// helpers:

func existOrCreateDir(path string) error {
	_, err := os.Stat(path)
	if errors.Is(err, os.ErrNotExist) {
		// if not existing, create
		return os.MkdirAll(path, 0755)
	}

	// if exists, no error
	return nil
}

func existsOrCreateFile(path string) error {
	_, err := os.Stat(path)
	if errors.Is(err, os.ErrNotExist) {
		// if no existing, create
		_, err = os.Create(path)
		return err
	}

	// if existst, error
	return fmt.Errorf("file %s already exists", path)
}

func getFileNameFromGooseOutput(output string) (string, error) {
	arr := strings.Split(output, " ")
	if len(arr) < 6 {
		return "", fmt.Errorf("goose output is malformed")
	}

	filename := arr[5]

	if !strings.HasPrefix(filename, config.MigrationsPath) {
		return "", fmt.Errorf("goose output is malformed")
	}

	return strings.TrimSuffix(filename, "\n"), nil
}
