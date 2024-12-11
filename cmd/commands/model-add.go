package commands

import (
	"fmt"

	"github.com/peterszarvas94/goat/cmd/helpers"
)

func ModelAdd(modelname string) error {
	// 0. CHECK -> TODO imporve

	if modelname == "" {
		return fmt.Errorf("Name can not be empty")
	}

	// 1. MIGRATIONS

	migrationFilepath, err := helpers.CreateMigrationFile(modelname, true)
	if err != nil {
		return err
	}

	fmt.Printf("Migration file created: %s\n", migrationFilepath)

	// 2. SCHEMA

	schemaFilePath, err := helpers.CreateSchemaFile(modelname, "")
	if err != nil {
		return err
	}

	fmt.Printf("Schema is created: %s\n", schemaFilePath)

	// 3. QUERIES

	queriesFilePath, err := helpers.CreateQueriesFile(modelname, "")
	if err != nil {
		return err
	}

	fmt.Printf("Queries are created: %s\n", queriesFilePath)

	// 4. GENREATE GO TYPES AND FUNCTIONS

	output, err := helpers.Cmd("sqlc", "generate")
	fmt.Println(output)
	if err != nil {
		return err
	}

	fmt.Println("Go functions and types are generated")

	fmt.Println("Run migration with \"goat migrate:up\"")

	return nil
}
