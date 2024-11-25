package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/peterszarvas94/goat/config"
)

func createQueriesFile(modelname, sql string) (string, error) {
	err := existOrCreateDir(config.QueriesDirPath)
	if err != nil {
		return "", err
	}

	queriesFilePath := filepath.Join(config.QueriesDirPath, fmt.Sprintf("%s.sql", modelname))
	err = existsOrCreateFile(queriesFilePath)
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
