package cmd

import (
	"fmt"

	"github.com/peterszarvas94/goat/pkg/utils"
	"github.com/spf13/cobra"
)

func createMigration(modelName string) error {
	migrationFilepath, err := utils.CreateMigrationFile(modelName, false)
	if err != nil {
		return err
	}

	fmt.Printf("Migration file created: %s\n", migrationFilepath)
	return nil
}

var migrationNewCmd = &cobra.Command{
	Use:                   "mig:new [title]",
	Aliases:               []string{"mn"},
	Short:                 "Add new empty migration file",
	Args:                  cobra.ExactArgs(1),
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		err := createMigration(args[0])
		if err != nil {
			fmt.Println(err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(migrationNewCmd)
}
