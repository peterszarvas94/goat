package config

import "path/filepath"

var (
	MigrationsDir     = filepath.Join("db", "migrations")
	SchemaDir         = filepath.Join("db", "sql", "schema")
	QueriesDir        = filepath.Join("db", "sql", "queries")
	ScriptsDir        = "scripts"
	StylesDir         = "styles"
	ImportMapFile     = filepath.Join(ScriptsDir, "importmap.json")
	TSConfigPahtsFile = filepath.Join(ScriptsDir, "tsconfig.paths.json")
	LogDir            = "logs"
)
