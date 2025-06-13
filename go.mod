module github.com/peterszarvas94/goat

go 1.24.1

require (
	github.com/a-h/templ v0.3.865
	github.com/joho/godotenv v1.5.1
	github.com/mattn/go-sqlite3 v1.14.28
	github.com/peterszarvas94/gohtml v0.0.1
	github.com/spf13/cobra v1.9.1
	github.com/yuin/goldmark v1.7.12
	github.com/yuin/goldmark-emoji v1.0.6
	go.abhg.dev/goldmark/frontmatter v0.2.0
	golang.org/x/crypto v0.38.0
)

require (
	github.com/BurntSushi/toml v1.2.1 // indirect
	golang.org/x/net v0.40.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

require (
	github.com/google/uuid v1.6.0
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/spf13/pflag v1.0.6 // indirect
)

retract [v0.0.0-0, v0.3.10]
