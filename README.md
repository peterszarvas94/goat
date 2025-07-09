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

### 1. Intsall

`go install github.com/peterszarvas94/goat@latest`

### 2. Initialize new project

`goat init [app] [--template basic-auth]?`

### 3. Go to the new project folder (if it is not the current folder)

`cd [app]`

### 4. Run dev server

`make dev`

### 5. Check out site with live-reload

http:/localhost:7331

## Development

### Clone the repository

```bash
git clone https://github.com/peterszarvas94/goat.git
cd goat
```

### Install dependencies

```bash
go mod tidy
```

### Build the CLI

```bash
make build
```

### Run tests

```bash
make test
```

### Create a release

```bash
# Interactive (prompts for version)
make release

# Or with version argument
make release-version VERSION=v1.2.3
```
