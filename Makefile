.PHONY: install tidy

install:
	go run ./install

tidy:
	go mod tidy
