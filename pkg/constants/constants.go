package constants

import "path/filepath"

// gen
var DoNotModify = "// Autogenerated by GOAT, do not delete or modify"

// db
var (
	DBPath        = "sqlite.db"
	DBDir         = "db"
	MigrationsDir = filepath.Join("db", "migrations")
	SQLDir        = filepath.Join("db", "sql")
	SchemaDir     = filepath.Join("db", "sql", "schema")
	QueriesDir    = filepath.Join("db", "sql", "queries")
)

// assets
var (
	AssetsDir         = "assets"
	CSSDir            = filepath.Join("assets", "css")
	CssPkgDir         = filepath.Join("assets", "css", "pkg")
	CssSrcDir         = filepath.Join("assets", "css", "src")
	JSDir             = filepath.Join("assets", "js")
	JsPkgDir          = filepath.Join("assets", "js", "pkg")
	JsSrcDir          = filepath.Join("assets", "js", "src")
	ImportMapFile     = filepath.Join("assets", "js", "importmap.json")
	TSConfigPahtsFile = filepath.Join("assets", "js", "tsconfig.paths.json")
	ContentDir        = "content"
	MarkdownDir       = filepath.Join("content", "md")
	NotFoundTemplate1 = filepath.Join("content", "md", "404.md")
	NotFoundTemplate2 = filepath.Join("content", "md", "404", "index.md")
	HtmlDir           = filepath.Join("content", "html")
	NotFoundFile      = filepath.Join("content", "html", "404", "index.html")
)

// examples
var (
	ExamplesDir = "examples"
	Examples    = []string{"bare", "basic-auth", "markdown"}
)

// log
var LogDir = "logs"

// versions
var (
	GooseVersion string = "v3.24.2"
	SqlcVersion  string = "v1.29.0"
	TemplVersion string = "v0.3.906"
)
