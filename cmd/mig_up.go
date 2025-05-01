package cmd

import (
	"fmt"
	"os"

	"github.com/peterszarvas94/goat/constants"
	"github.com/peterszarvas94/goat/utils"
	"github.com/spf13/cobra"
)

func migrate(direction string) error {
	dbPath := os.Getenv("DBPATH")
	if dbPath == "" {
		return fmt.Errorf("DBPATH is missing from PATH")
	}

	err := utils.CreateDirIfNotExists(constants.MigrationsDir)
	if err != nil {
		return err
	}

	err = utils.Cmd("goose", "-dir", constants.MigrationsDir, "sqlite3", dbPath, direction)
	if err != nil {
		return err
	}

	return nil
}

var migrateUpCmd = &cobra.Command{
	Use:                   "mig:up",
	Short:                 "Run up migrations",
	Args:                  cobra.ExactArgs(0),
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		err := migrate("up")
		if err != nil {
			fmt.Println(err.Error())
		}
	},
}
