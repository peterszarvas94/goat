.PHONY: publish install

# e.g. `make run publish`
publish:
	go run ./publish

install:
	go install ./...
