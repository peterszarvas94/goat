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

	output, err := helpers.Cmd("goose", "-dir", config.MigrationsPath, "sqlite3", dbPath, direction)
	fmt.Println(output)
	if err != nil {
		return err
	}

	return nil
}

func migrateUpInitial() error {
	output, err := helpers.Cmd("goose", "-dir", config.MigrationsPath, "sqlite3", "sqlite.db", "up")
	fmt.Println(output)
	if err != nil {
		return err
	}

	return nil
}
