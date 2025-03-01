package config

import "path/filepath"

var (
	MigrationsDir     = filepath.Join("db", "migrations")
	SchemaDir         = filepath.Join("db", "sql", "schema")
	QueriesDir        = filepath.Join("db", "sql", "queries")
	ScriptsDir        = "scripts"
	StylesDir         = "styles"
	ImportMapFile     = "importmap.json"
	TSConfigPahtsFile = "tsconfig.paths.json"
	LogDir            = "logs"
)
