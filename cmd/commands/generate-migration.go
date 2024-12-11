package commands

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/peterszarvas94/goat/cmd/helpers"
	"github.com/peterszarvas94/goat/config"
)

func GenerateMigration(statement string, table string) (string, error) {
	dbPath := os.Getenv("DBPATH")
	if dbPath == "" {
		return "", errors.New("DBPATH is missing from PATH")
	}

	oldSchema, err := helpers.Cmd("sqlite3", dbPath, fmt.Sprintf(".schema %s", table))
	if err != nil {
		return "", err
	}

	db1, err := os.Create("db1.db")
	if err != nil {
		return "", err
	}

	output, err := helpers.Cmd("sqlite3", db1.Name(), fmt.Sprintf("%s", oldSchema))
	fmt.Println(output)
	if err != nil {
		return "", err
	}

	db2, err := os.Create("db2.db")
	if err != nil {
		return "", err
	}

	schemaName := fmt.Sprintf("%s.sql", table)
	schemaPath := filepath.Join(config.SchemaDirPath, schemaName)
	newSchema, err := helpers.ExistsOrCreateFile(schemaPath)
	if err != nil {
		return "", err
	}

	if newSchema == "" {
		newSchema = helpers.GetDefaultSchemaSql(table)
	}

	output, err = helpers.Cmd("sqlite3", db2.Name(), fmt.Sprintf("%s", newSchema))
	fmt.Println(output)
	if err != nil {
		return "", err
	}

	up, err := helpers.Cmd("atlas", "schema", "diff",
		"--from", fmt.Sprintf("sqlite3://%s", db1.Name()),
		"--to", fmt.Sprintf("sqlite3://%s", db2.Name()))
	fmt.Println(up)
	if err != nil {
		return "", err
	}

	down, err := helpers.Cmd("atlas", "schema", "diff",
		"--from", fmt.Sprintf("sqlite3://%s", db2.Name()),
		"--to", fmt.Sprintf("sqlite3://%s", db1.Name()))
	fmt.Println(down)
	if err != nil {
		return "", err
	}

	err = helpers.ExistsOrCreateDir(config.MigrationsPath)
	if err != nil {
		return "", err
	}

	output, err = helpers.Cmd("goose", "-dir", config.MigrationsPath, "create", fmt.Sprintf("%s_%s", statement, table), "sql")
	fmt.Println(output)
	if err != nil {
		return "", err
	}

	migrationFilepath, err := helpers.GetFileNameFromGooseOutput(output)
	if err != nil {
		return "", err
	}

	migration := fmt.Sprintf(`-- +goose Up
%s

-- +goose Down
%s
		`, up, down)

	err = os.WriteFile(migrationFilepath, []byte(migration), 0644)
	if err != nil {
		return "", err
	}

	err = os.Remove("db1.db")
	if err != nil {
		return "", err
	}

	err = os.Remove("db2.db")
	if err != nil {
		return "", err
	}

	return migrationFilepath, nil
}
