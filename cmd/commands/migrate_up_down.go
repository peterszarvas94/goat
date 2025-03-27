package commands

import (
	"fmt"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/peterszarvas94/goat/cmd/helpers"
	"github.com/peterszarvas94/goat/constants"
)

func MigrateUpDown(direction string) error {
	dbPath := os.Getenv("DBPATH")
	if dbPath == "" {
		return fmt.Errorf("DBPATH is missing from PATH")
	}

	err := helpers.ExistsOrCreateDir(constants.MigrationsDir)
	if err != nil {
		return err
	}

	output, err := helpers.Cmd("goose", "-dir", constants.MigrationsDir, "sqlite3", dbPath, direction)
	fmt.Println(output)
	if err != nil {
		return err
	}

	return nil
}

func migrateUpInitial() error {
	err := helpers.ExistsOrCreateDir(constants.MigrationsDir)
	if err != nil {
		return err
	}

	entries, err := os.ReadDir(constants.MigrationsDir)
	if err != nil {
		return err
	}

	// no migrations found
	if len(entries) == 0 {
		return nil
	}

	output, err := helpers.Cmd("goose", "-dir", constants.MigrationsDir, "sqlite3", constants.DBPath, "up")
	fmt.Println(output)
	if err != nil {
		return err
	}

	fmt.Println("Default db schema is migrated")

	return err
}
