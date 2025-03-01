package helpers

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/peterszarvas94/goat/config"
)

func CreateSchemaFile(modelname string, sql string) (string, error) {
	err := ExistsOrCreateDir(config.SchemaDir)
	if err != nil {
		return "", err
	}

	schemaFilePath := filepath.Join(config.SchemaDir, fmt.Sprintf("%s.sql", modelname))
	err = createFileIfNotExists(schemaFilePath)
	if err != nil {
		return "", err
	}

	modelSQL := sql
	if modelSQL == "" {
		modelSQL = GetDefaultSchemaSql(modelname)
	}

	err = os.WriteFile(schemaFilePath, []byte(modelSQL), 0644)
	if err != nil {
		return "", err
	}

	return schemaFilePath, nil
}
