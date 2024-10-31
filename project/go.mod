module project

go 1.22.5

require github.com/a-h/templ v0.2.778

require (
	github.com/joho/godotenv v1.5.1
	github.com/peterszarvas94/goat v0.2.778
)

require (
	github.com/antlr4-go/antlr/v4 v4.13.0 // indirect
	github.com/coder/websocket v1.8.12 // indirect
	github.com/tursodatabase/libsql-client-go v0.0.0-20240902231107-85af5b9d094d // indirect
	golang.org/x/exp v0.0.0-20240325151524-a685a6edb6d8 // indirect
)

replace github.com/peterszarvas94/goat => ../lib
