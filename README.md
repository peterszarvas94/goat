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

## CLI Usage

### 1. Intsall

```bash
go install github.com/peterszarvas94/goat@latest
```

### 2. Initialize new project

```bash
goat init [app] [--template basic-auth]?
```

### 3. Go to the new project folder (if it is not the current folder)

```bash
cd [app]
```

### 4. Run dev server

```bash
make dev
```

### 5. Check out site with live-reload

[http://localhost:7331](http://localhost:7331)

## CLI Development

### Clone the repository

```bash
git clone https://github.com/peterszarvas94/goat.git
cd goat
```

### Install dependencies for CLI

```bash
go mod tidy
```

### Build the CLI

```bash
go build -o tmp/goat ./main.go
```

### Install the CLI

```bash
go install ./...
```

### Run tests

```bash
go test ./...
```

### Update templ files

```bash
go run ./scripts/templ-update
```

### Create a release

```bash
./scripts/release.sh v1.2.3
```

### Test release locally (using goreleaser)

```bash
# Test release configuration
goreleaser release --snapshot --clean

# Build snapshot locally
goreleaser build --snapshot --clean
```

## Example Development

The examples in the `examples/` directory demonstrate how to use the GOAT framework. Each example is a complete application with its own `go.mod` file.

### Setup workspace for development

```bash
# Create workspace to use local goat code in examples
go work init . examples/bare examples/basic-auth examples/markdown
```

### Working with examples

```bash
# Navigate to any example
cd examples/bare

# Install dependencies and run
go mod tidy
make dev

# Or run specific commands
make build
make test
```

### Example structure

Each example contains:

- `go.mod` - Module definition (references published goat version)
- `Makefile` - Development commands (`make dev`, `make build`, etc.)
- Standard GOAT application structure (cmd/, views/, controllers/, etc.)

### Development workflow

1. **Local development**: Examples use workspace to reference local goat code
2. **Release**: Examples are updated to reference the new published version
3. **User experience**: Users get examples that work with the published version

### Tidy all modules

```bash
# Tidy root module only
go mod tidy

# Tidy all examples individually
for dir in examples/*/; do (cd "$dir" && go mod tidy); done
```
