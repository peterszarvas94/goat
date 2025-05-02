package cmd

import (
	"fmt"

	"github.com/peterszarvas94/goat/pkg/utils"
	"github.com/spf13/cobra"
)

func addModel(modelname string) error {
	/* 0. CHECK */
	// TODO: improve

	if modelname == "" {
		return fmt.Errorf("Name can not be empty")
	}

	/* 1. MIGRATIONS */

	migrationFilepath, err := utils.CreateMigrationFile(modelname, true)
	if err != nil {
		return err
	}

	fmt.Printf("Migration file created: %s\n", migrationFilepath)

	/* 2. SCHEMA */

	schemaFilePath, err := utils.CreateSchemaFile(modelname, "")
	if err != nil {
		return err
	}

	fmt.Printf("Schema is created: %s\n", schemaFilePath)

	/* 3. QUERIES */

	queriesFilePath, err := utils.CreateQueriesFile(modelname)
	if err != nil {
		return err
	}

	fmt.Printf("Queries are created: %s\n", queriesFilePath)

	/* 4. SQLC */

	err = utils.Cmd("sqlc", "generate")
	if err != nil {
		return err
	}

	fmt.Println("Go functions and types are generated")

	fmt.Println("Run migration with \"goat mig:up\"")

	return nil
}

var modelAddCmd = &cobra.Command{
	Use:                   "model:add [name]",
	Short:                 "Add new model",
	Args:                  cobra.ExactArgs(1),
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		err := addModel(args[0])
		if err != nil {
			fmt.Println(err.Error())
		}
	},
}
