.PHONY: publish install templ-update

# e.g. `make run publish`
publish:
	go run ./publish

templ-update:
	go run ./templ-update

install:
	go install ./...
