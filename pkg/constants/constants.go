package constants

import "path/filepath"

// gen
var DoNotModify = "// Autogenerated by GOAT, do not delete or modify"

// db
var (
	DBPath        = "sqlite.db"
	DBDir         = "db"
	MigrationsDir = filepath.Join(DBDir, "migrations")
	SQLDir        = filepath.Join(DBDir, "sql")
	SchemaDir     = filepath.Join(SQLDir, "schema")
	QueriesDir    = filepath.Join(SQLDir, "queries")
)

// assets
var (
	AssetsDir         = "assets"
	CSSDir            = filepath.Join(AssetsDir, "css")
	UserStylesDir     = filepath.Join(CSSDir, "src")
	JSDir             = filepath.Join(AssetsDir, "js")
	UserScriptsDir    = filepath.Join(JSDir, "src")
	ImportMapFile     = filepath.Join(JSDir, "importmap.json")
	TSConfigPahtsFile = filepath.Join(JSDir, "tsconfig.paths.json")
)

// log
var LogDir = "logs"
