package utils

import (
	"fmt"
	"os"

	"github.com/peterszarvas94/goat/constants"
)

func CreateMigrationFile(modelname string, create bool) (string, error) {
	err := CreateDirIfNotExists(constants.MigrationsDir)
	if err != nil {
		return "", err
	}

	output, err := Cmd("goose", "-dir", constants.MigrationsDir, "create", fmt.Sprintf("create_%s_table", modelname), "sql")
	if err != nil {
		return "", err
	}

	filePath, err := getFileNameFromGooseOutput(output)
	if err != nil {
		return "", err
	}

	sql := generateMigrationSql(modelname, create)

	err = os.WriteFile(filePath, []byte(sql), 0644)
	if err != nil {
		return "", err
	}

	return filePath, nil
}
