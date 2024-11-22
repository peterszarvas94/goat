run:
	go run cmd/server/main.go

dump:
	sqlite3 project/sqlite.db .dump > ./dump.sql

get:
	go get github.com/peterszarvas94/goat
	go mod tidy

build:
	go build -o ./main cmd/server/main.go
