# Load environment variables from .env file
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
dev/css:
	@echo "Generating css"
	@npm run css:watch -C assets/daisyui

dev/templ:
	@sleep 1
	@echo "Generating temmpaltes"
	@templ generate --watch --proxy="http://localhost:$(PORT)" --open-browser=false 

dev/server:
	@sleep 2
	@echo "Running server"
	air -c .air.server.toml

dev/assets:
	@sleep 3
	@echo "Watching assets"
	air -c .air.assets.toml

dev:
	@rm -rf tmp && make -j4 dev/css dev/templ dev/server dev/assets

# dump db
dump:
	sqlite3 "$(DBPATH)" .dump > ./dump.sql

# build binary
build:
	go build -o tmp/main cmd/main.go

# run binary
run:
	tmp/main

.PHONY: dev/css dev/templ dev/server dev/assets dev dump build run
