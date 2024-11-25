package config

import "path/filepath"

var (
	MigrationsPath = filepath.Join("db", "migrations")
	SchemaDirPath  = filepath.Join("db", "sql", "schema")
	QueriesDirPath = filepath.Join("db", "sql", "queries")
)
