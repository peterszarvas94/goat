# AGENTS.md - GOAT Framework Development Guide

## Build/Test Commands
- `go test ./...` - Run all tests
- `go test ./pkg/content` - Run single package tests  
- `make build` - Build binary to tmp/main
- `make dev` - Start development server with live reload
- `go install ./...` - Install CLI tool
- `make release` - Create release (interactive, prompts for version)
- `make release-version VERSION=v1.2.3` - Create release with specific version

## Release Process
- **Release script**: `./scripts/release.sh v1.2.3` handles the full release process
- **Automated CI**: GitHub Actions + GoReleaser builds and publishes releases
- **Example dependencies**: Release script updates example go.mod files to reference new version
- **Workspace setup**: CI generates `go.work` dynamically for testing with local dependencies
- **Clean checksums**: go.sum files are gitignored to avoid checksum conflicts

## Development Setup
- **Workspace**: Run `go work init . examples/bare examples/basic-auth examples/markdown` for local development
- **Examples**: Use local goat code during development via workspace
- **Testing**: Full test suite runs in CI with workspace dependency resolution

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

## File Structure Notes
- **examples/**: Contains example applications that demonstrate framework usage
- **examples/*/go.sum**: Gitignored, generated locally but not committed
- **go.work**: Gitignored, generated in CI and locally for development
- **scripts/release.sh**: Automated release script that handles version updates and tagging
- **.github/workflows/release.yml**: CI pipeline for automated releases