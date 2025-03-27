package helpers

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/peterszarvas94/goat/constants"
)

func CreateQueriesFile(modelname, sql string) (string, error) {
	err := ExistsOrCreateDir(constants.QueriesDir)
	if err != nil {
		return "", err
	}

	queriesFilePath := filepath.Join(constants.QueriesDir, fmt.Sprintf("%s.sql", modelname))
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
