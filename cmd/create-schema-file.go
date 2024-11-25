package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/peterszarvas94/goat/config"
)

func createSchemaFile(modelname string, sql string) (string, error) {
	err := existOrCreateDir(config.SchemaDirPath)
	if err != nil {
		return "", err
	}

	schemaFilePath := filepath.Join(config.SchemaDirPath, fmt.Sprintf("%s.sql", modelname))
	err = existsOrCreateFile(schemaFilePath)
	if err != nil {
		return "", err
	}

	modelSQL := sql
	if modelSQL == "" {
		modelSQL = getDefaultSchemaSql(modelname)
	}

	err = os.WriteFile(schemaFilePath, []byte(modelSQL), 0644)
	if err != nil {
		return "", err
	}

	return schemaFilePath, nil
}
