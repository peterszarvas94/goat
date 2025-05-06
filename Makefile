.PHONY: install publish init-dev 

install:
	go run ./scripts/install/main.go

publish:
	go run ./scripts/publish/main.go

init-dev:
	go run ./scripts/init-dev/main.go
