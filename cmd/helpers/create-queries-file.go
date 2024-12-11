package helpers

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/peterszarvas94/goat/config"
)

func CreateQueriesFile(modelname, sql string) (string, error) {
	err := ExistsOrCreateDir(config.QueriesDirPath)
	if err != nil {
		return "", err
	}

	queriesFilePath := filepath.Join(config.QueriesDirPath, fmt.Sprintf("%s.sql", modelname))
	err = createFileIfNotExists(queriesFilePath)
	if err != nil {
		return "", err
	}

	queriesSQL := sql
	if queriesSQL == "" {
		queriesSQL = getDefaultQueriesSql(modelname)
	}

	err = os.WriteFile(queriesFilePath, []byte(queriesSQL), 0644)
	if err != nil {
		return "", err
	}

	return queriesFilePath, nil
}
