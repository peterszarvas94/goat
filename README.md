# GOAT

GO Application Toolkit

The all-in-one web framework for go (in progress!)

## Dependencies

### You need to install

- make: [https://www.gnu.org/software/make](https://www.gnu.org/software/make)
- sqlite3: [https://www.sqlite.org/cli.html](https://www.sqlite.org/cli.html)

### GOAT will install for you

- air: [https://github.com/air-verse/air](https://github.com/air-verse/air)
- templ: [https://github.com/a-h/templ](https://github.com/a-h/templ)
- sqlc: [https://github.com/sqlc-dev/sqlc](https://github.com/sqlc-dev/sqlc)

## CLI

### Quick Start

```bash
# Install CLI
go install github.com/peterszarvas94/goat@latest

# Create new project
goat init myapp --template basic-auth
cd myapp

# Start development server
make dev
```

Visit [http://localhost:7331](http://localhost:7331)

## CLI Development

```bash
# Clone and setup
git clone https://github.com/peterszarvas94/goat.git
cd goat
go mod tidy
```

### Common Commands

```bash
go build -o tmp/goat ./main.go        # Build CLI
go install ./...                      # Install CLI globally
go test ./...                         # Run tests
go run ./scripts/templ-update         # Update templ files
./scripts/release.sh v1.2.3           # Create release
goreleaser release --snapshot --clean # Test release locally
```

## Example Development

The examples in the `examples/` directory demonstrate how to use the GOAT framework. Each example is a complete application with its own `go.mod` file.

### Setup & Usage

```bash
# Setup workspace (links examples to local goat code)
go work init . examples/bare examples/basic-auth examples/markdown

# Work with any example
cd examples/bare
go mod tidy && make dev
```

### Notes

- Examples use workspace during development (local goat code)
- Examples reference published versions for end users
- Each example has its own `Makefile` with `make dev`, `make build`, etc.

```bash
go mod tidy                                        # Tidy root only
for dir in examples/*/; do (cd "$dir" && go mod tidy); done  # Tidy all examples
```