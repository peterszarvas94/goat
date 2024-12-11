package commands

import (
	"fmt"

	"github.com/peterszarvas94/goat/cmd/helpers"
)

func NewMigration(modelName string) error {
	migrationFilepath, err := helpers.CreateMigrationFile(modelName, false)
	if err != nil {
		return err
	}

	fmt.Printf("Migration file created: %s\n", migrationFilepath)
	return nil
}
