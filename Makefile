.PHONY: publish install templ-update build test release-test release-local

# Create and push release tag
publish:
	@read -p "Enter version (v1.2.3): " version; \
	./scripts/release.sh $$version

templ-update:
	go run ./templ-update

install:
	go install ./...

build:
	go build -o tmp/goat ./main.go

test:
	go test ./...

# Test goreleaser configuration locally
release-test:
	goreleaser release --snapshot --clean

# Build snapshot locally
release-local:
	goreleaser build --snapshot --clean
