package utils

import (
	"fmt"
	"path/filepath"

	"github.com/peterszarvas94/goat/constants"
)

func CreateSchemaFile(modelname, sql string) (string, error) {
	err := CreateDirIfNotExists(constants.SchemaDir)
	if err != nil {
		return "", err
	}

	schemaFilePath := filepath.Join(constants.SchemaDir, fmt.Sprintf("%s.sql", modelname))
	file, err := CreateNonExistingFile(schemaFilePath)
	if err != nil {
		return "", err
	}

	modelSQL := sql
	if modelSQL == "" {
		modelSQL = generateSchemaSql(modelname)
	}

	_, err = file.Write([]byte(modelSQL))
	if err != nil {
		return "", err
	}

	return schemaFilePath, nil
}
