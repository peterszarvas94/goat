#a Load environment variables from .env file
ifneq (,$(wildcard ./.env))
	include .env
	export
endif

ifndef DBPATH
  $(error DBPATH is not set in the environment)
endif

ifndef PORT
  $(error PORT is not set in the environment)
endif

ifndef GOATENV 
  $(error GOATENV is not set in the environment)
endif

# dev server
dev/templ:
	@echo "Generating temmpaltes"
	@templ generate --watch --proxy="http://localhost:$(PORT)" --open-browser=false 

dev/server:
	@sleep 1
	@echo "Running server"
	air -c .air.server.toml

dev/assets:
	@sleep 2
	@echo "Watching assets"
	air -c .air.assets.toml

dev:
	@rm -rf tmp && make -j3 dev/templ dev/server dev/assets

# dump db
dump:
	sqlite3 "$(DBPATH)" .dump > ./dump.sql

# build binary
build:
	go build -o tmp/main cmd/main.go

# run binary
run:
	tmp/main
