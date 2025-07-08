# AGENTS.md - GOAT Framework Development Guide

## Build/Test Commands
- `go test ./...` - Run all tests
- `go test ./pkg/content` - Run single package tests  
- `make build` - Build binary to tmp/main
- `make dev` - Start development server with live reload
- `go install ./...` - Install CLI tool
- `make publish` - Publish/release

## Code Style & Conventions
- **Imports**: Standard library first, then external packages, then local packages with blank lines between groups
- **Naming**: Use camelCase for variables/functions, PascalCase for exported types, snake_case for file names
- **Types**: Define structs with clear field names and comments for exported types
- **Error Handling**: Always check errors, wrap with context using fmt.Errorf("description: %w", err)
- **Comments**: Document exported functions/types, use // for single line, avoid obvious comments
- **Testing**: Use table-driven tests with descriptive test names, place in *_test.go files
- **Packages**: Keep packages focused, use pkg/ for reusable components, cmd/ for executables
- **File Structure**: Group related functionality, use init() for package setup
- **Formatting**: Use tabs for indentation, gofmt standard formatting
- **Constants**: Define in constants/ package, use ALL_CAPS for package-level constants

## Framework Notes
- Uses templ for HTML templating, goldmark for markdown processing
- Built-in dev server with air for live reload
- SQLite database with sqlc for type-safe queries
- CLI built with cobra, uses make for build automation