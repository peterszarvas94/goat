package cmd

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func modelAdd(name string) error {
	if name == "" {
		return fmt.Errorf("Name can not be empty")
	}

	schemaDirPath := filepath.Join("sql", "schema")
	err := existOrCreateDir(schemaDirPath)
	if err != nil {
		return err
	}

	queriesDirPath := filepath.Join("sql", "queries")
	err = existOrCreateDir(queriesDirPath)
	if err != nil {
		return err
	}

	schemaFilePath := filepath.Join(schemaDirPath, fmt.Sprintf("%s.sql", name))
	err = existsOrCreateFile(schemaFilePath)
	if err != nil {
		return err
	}

	modelSQL := fmt.Sprintf(`CREATE TABLE %s (
    id TEXT PRIMARY KEY,
    name TEXT NOT NULL
);`,
		name,
	)
	err = os.WriteFile(schemaFilePath, []byte(modelSQL), 0644)
	if err != nil {
		return err
	}

	fmt.Printf("Schema is created: %s\n", schemaFilePath)

	queriesFilePath := filepath.Join(queriesDirPath, fmt.Sprintf("%s.sql", name))
	err = existsOrCreateFile(queriesFilePath)
	if err != nil {
		return err
	}

	upperCaseName := fmt.Sprintf("%s%s", strings.ToUpper(string(name[0])), name[1:])

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
		upperCaseName,
		name,
		upperCaseName,
		name,
		upperCaseName,
		name,
		upperCaseName,
		name,
		upperCaseName,
		name,
	)

	err = os.WriteFile(queriesFilePath, []byte(queriesSQL), 0644)
	if err != nil {
		return err
	}

	fmt.Printf("Queries are created: %s\n", queriesFilePath)

	return nil
}

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
