package cmd

import (
	"fmt"
)

func modelAdd(modelname string, showHelperText bool) error {
	// 0. CHECK -> TODO imporve

	if modelname == "" {
		return fmt.Errorf("Name can not be empty")
	}

	// 1. MIGRATIONS

	migrationFilepath, err := createMigrationFile(modelname, "")
	if err != nil {
		return err
	}

	fmt.Printf("Migration file created: %s\n", migrationFilepath)

	// 2. SCHEMA

	schemaFilePath, err := createSchemaFile(modelname, "")
	if err != nil {
		return err
	}

	fmt.Printf("Schema is created: %s\n", schemaFilePath)

	// 3. QUERIES

	queriesFilePath, err := createQueriesFile(modelname, "")
	if err != nil {
		return err
	}

	fmt.Printf("Queries are created: %s\n", queriesFilePath)

	// 4. GENREATE GO TYPES AND FUNCTIONS

	output, err := cmd("sqlc", "generate")
	fmt.Println(output)
	if err != nil {
		return err
	}

	fmt.Println("Go functions and types are generated")

	if showHelperText {
		fmt.Println("Run migration with \"goat migrate:up\"")
	}

	return nil
}
