package utils

import (
	"fmt"
	"path/filepath"

	"github.com/peterszarvas94/goat/constants"
)

func CreateQueriesFile(modelname string) (string, error) {
	err := CreateDirIfNotExists(constants.QueriesDir)
	if err != nil {
		return "", err
	}

	filePath := filepath.Join(constants.QueriesDir, fmt.Sprintf("%s.sql", modelname))

	file, err := CreateNonExistingFile(filePath)
	if err != nil {
		return "", err
	}

	sql := generateQueriesSql(modelname)

	_, err = file.Write([]byte(sql))
	if err != nil {
		return "", err
	}

	return filePath, nil
}
