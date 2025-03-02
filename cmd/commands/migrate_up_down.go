package commands

import (
	"fmt"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/peterszarvas94/goat/cmd/helpers"
	"github.com/peterszarvas94/goat/config"
)

func MigrateUpDown(direction string) error {
	dbPath := os.Getenv("DBPATH")
	if dbPath == "" {
		return fmt.Errorf("DBPATH is missing from PATH")
	}

	err := helpers.ExistsOrCreateDir(config.MigrationsDir)
	if err != nil {
		return err
	}

	output, err := helpers.Cmd("goose", "-dir", config.MigrationsDir, "sqlite3", dbPath, direction)
	fmt.Println(output)
	if err != nil {
		return err
	}

	return nil
}

func migrateUpInitial() error {
	err := helpers.ExistsOrCreateDir(config.MigrationsDir)
	if err != nil {
		return err
	}

	entries, err := os.ReadDir(config.MigrationsDir)
	if err != nil {
		return err
	}

	// no migrations found
	if len(entries) == 0 {
		return nil
	}

	output, err := helpers.Cmd("goose", "-dir", config.MigrationsDir, "sqlite3", config.DBPath, "up")
	fmt.Println(output)
	if err != nil {
		return err
	}

	fmt.Println("Default db schema is migrated")

	return err
}
