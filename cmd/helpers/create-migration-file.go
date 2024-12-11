package helpers

import (
	"fmt"
	"os"

	"github.com/peterszarvas94/goat/config"
)

func CreateMigrationFile(modelname string, withSql bool) (string, error) {
	err := ExistsOrCreateDir(config.MigrationsPath)
	if err != nil {
		return "", err
	}

	output, err := Cmd("goose", "-dir", config.MigrationsPath, "create", fmt.Sprintf("create_%s_table", modelname), "sql")
	fmt.Println(output)
	if err != nil {
		return "", err
	}

	migrationFilepath, err := GetFileNameFromGooseOutput(output)
	if err != nil {
		return "", err
	}

	migrationSQL := ""
	if !withSql {
		migrationSQL = `-- +goose Up

-- +goose Down`
	} else {
		migrationSQL = getDefaultMigraionSql(modelname)
	}

	err = os.WriteFile(migrationFilepath, []byte(migrationSQL), 0644)
	if err != nil {
		return "", err
	}

	return migrationFilepath, nil
}
