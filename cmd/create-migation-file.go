package cmd

import (
	"fmt"
	"os"

	"github.com/peterszarvas94/goat/config"
)

func createMigrationFile(modelname, sql string) (string, error) {
	err := existOrCreateDir(config.MigrationsPath)
	if err != nil {
		return "", err
	}

	output, err := cmd("goose", "-dir", config.MigrationsPath, "create", fmt.Sprintf("create_%s_table", modelname), "sql")
	fmt.Println(output)
	if err != nil {
		return "", err
	}

	migrationFilepath, err := getFileNameFromGooseOutput(output)
	if err != nil {
		return "", err
	}

	migrationSQL := sql
	if migrationSQL == "" {
		migrationSQL = getDefaultMigraionSql(modelname)
	}

	err = os.WriteFile(migrationFilepath, []byte(migrationSQL), 0644)
	if err != nil {
		return "", err
	}

	return migrationFilepath, nil
}
