module github.com/peterszarvas94/goat

go 1.24.1

require (
	github.com/a-h/templ v0.3.865
	github.com/joho/godotenv v1.5.1
	github.com/mattn/go-sqlite3 v1.14.28
	github.com/spf13/cobra v1.9.1
	golang.org/x/crypto v0.38.0
)

require (
	github.com/google/uuid v1.6.0
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/spf13/pflag v1.0.6 // indirect
)

retract [v0.0.0-0, v0.3.10]
