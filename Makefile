run:
	cd project && \
	templ generate && \
	go run cmd/server/main.go

tidy:
	cd lib && \
	go mod tidy;
	cd project && \
	go mod tidy;

dump:
	sqlite3 project/sqlite.db .dump > ./dump.sql
