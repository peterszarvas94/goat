package config

import "path/filepath"

// db
var (
	DBDir         = "db"
	MigrationsDir = filepath.Join(DBDir, "migrations")
	SQLDir        = filepath.Join(DBDir, "sql")
	SchemaDir     = filepath.Join(SQLDir, "schema")
	QueriesDir    = filepath.Join(SQLDir, "queries")
)

// assets
var AssetsDir = "assets"

// css
var (
	CSSDir        = filepath.Join(AssetsDir, "css")
	UserStylesDir = filepath.Join(CSSDir, "src")
)

// js
var (
	JSDir             = filepath.Join(AssetsDir, "js")
	UserScriptsDir    = filepath.Join(JSDir, "src")
	ImportMapFile     = filepath.Join(JSDir, "importmap.json")
	TSConfigPahtsFile = filepath.Join(JSDir, "tsconfig.paths.json")
)

// log
var LogDir = "logs"
