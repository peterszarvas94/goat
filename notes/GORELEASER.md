# GoReleaser Integration Guide for GOAT Framework

## Overview

GoReleaser is a release automation tool that can replace your current manual publish script with automated cross-platform builds, GitHub releases, and proper semantic versioning.

## Current vs GoReleaser Approach

### Current Publish Script Issues
- Manual version management in `pkg/version/version.go`
- Complex go.mod manipulation for examples
- No cross-platform builds
- No GitHub releases or changelog generation
- Manual git operations prone to errors
- No rollback mechanism

### GoReleaser Benefits
- Automated cross-platform builds (Linux, macOS, Windows)
- GitHub releases with changelogs
- Homebrew tap integration
- Docker image publishing
- Checksums and signing
- Rollback capabilities

## Installation

```bash
# Install GoReleaser
brew install goreleaser/tap/goreleaser

# Or with Go
go install github.com/goreleaser/goreleaser@latest
```

## Configuration

Create `.goreleaser.yaml` in your project root:

```yaml
version: 1

before:
  hooks:
    - go mod tidy
    - go generate ./...

builds:
  - main: ./main.go
    binary: goat
    env:
      - CGO_ENABLED=0
    goos:
      - linux
      - windows
      - darwin
    goarch:
      - amd64
      - arm64
    ldflags:
      - -s -w -X github.com/peterszarvas94/goat/pkg/version.Version={{.Version}}

archives:
  - format: tar.gz
    name_template: >-
      {{ .ProjectName }}_
      {{- title .Os }}_
      {{- if eq .Arch "amd64" }}x86_64
      {{- else if eq .Arch "386" }}i386
      {{- else }}{{ .Arch }}{{ end }}
      {{- if .Arm }}v{{ .Arm }}{{ end }}
    format_overrides:
      - goos: windows
        format: zip

checksum:
  name_template: 'checksums.txt'

snapshot:
  name_template: "{{ incpatch .Version }}-next"

changelog:
  sort: asc
  filters:
    exclude:
      - '^docs:'
      - '^test:'

release:
  github:
    owner: peterszarvas94
    name: goat
  draft: false
  prerelease: auto

brews:
  - name: goat
    homepage: https://github.com/peterszarvas94/goat
    description: "Go web framework with templ and HTMX"
    repository:
      owner: peterszarvas94
      name: homebrew-tap
    folder: Formula
    install: |
      bin.install "goat"
```

## Workflow Integration

### GitHub Actions (Recommended)

Create `.github/workflows/release.yml`:

```yaml
name: Release

on:
  push:
    tags:
      - 'v*'

permissions:
  contents: write

jobs:
  goreleaser:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
        with:
          fetch-depth: 0
      - uses: actions/setup-go@v4
        with:
          go-version: stable
      - uses: goreleaser/goreleaser-action@v5
        with:
          distribution: goreleaser
          version: latest
          args: release --clean
        env:
          GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
```

## Migration Steps

### 1. Update Version Management

Replace manual version in `pkg/version/version.go`:

```go
package version

// Version is set by GoReleaser during build
var Version = "dev"
```

### 2. Create Release Script

Replace `publish/main.go` with simpler `scripts/release.sh`:

```bash
#!/bin/bash
set -e

VERSION=$1
if [ -z "$VERSION" ]; then
    echo "Usage: $0 <version>"
    echo "Example: $0 v1.2.3"
    exit 1
fi

# Validate version format
if [[ ! $VERSION =~ ^v[0-9]+\.[0-9]+\.[0-9]+$ ]]; then
    echo "Version must be in format v1.2.3"
    exit 1
fi

# Check for uncommitted changes
if ! git diff-index --quiet HEAD --; then
    echo "You have uncommitted changes"
    exit 1
fi

# Update examples go.mod files (keep your existing logic)
echo "Updating example go.mod files..."
# ... your existing go.mod update logic ...

# Create and push tag
git add .
git commit -m "chore: prepare release $VERSION"
git tag $VERSION
git push origin main
git push origin $VERSION

echo "Release $VERSION created. GitHub Actions will handle the rest."
```

### 3. Update Makefile

```makefile
.PHONY: release build-snapshot

release:
	@read -p "Enter version (v1.2.3): " version; \
	./scripts/release.sh $$version

build-snapshot:
	goreleaser build --snapshot --clean

test-release:
	goreleaser release --snapshot --clean
```

## Usage

### Development Builds
```bash
# Test release locally
make test-release

# Build snapshot
make build-snapshot
```

### Production Release
```bash
# Create release
make release
# Enter: v1.2.3

# Or directly
./scripts/release.sh v1.2.3
```

## Advanced Features

### Docker Integration
Add to `.goreleaser.yaml`:

```yaml
dockers:
  - image_templates:
      - "peterszarvas94/goat:{{ .Tag }}"
      - "peterszarvas94/goat:latest"
    dockerfile: Dockerfile
```

### Homebrew Tap
GoReleaser can automatically update your Homebrew formula in a separate repository.

### Signing and Notarization
For macOS notarization and code signing:

```yaml
signs:
  - artifacts: checksum
    args:
      - "--batch"
      - "--local-user"
      - "{{ .Env.GPG_FINGERPRINT }}"
      - "--output"
      - "${signature}"
      - "--detach-sign"
      - "${artifact}"
```

## Benefits Over Current Script

1. **Cross-platform builds**: Automatic builds for multiple OS/arch combinations
2. **GitHub integration**: Automatic releases with changelogs
3. **Package management**: Homebrew, Scoop, etc.
4. **Checksums**: Automatic generation for security
5. **Rollback**: Easy to revert releases
6. **CI/CD ready**: Integrates with GitHub Actions
7. **Less maintenance**: No custom git/version logic

## Migration Timeline

1. **Phase 1**: Add GoReleaser config, test locally
2. **Phase 2**: Set up GitHub Actions workflow
3. **Phase 3**: Replace publish script with release script
4. **Phase 4**: Add advanced features (Docker, Homebrew)

This approach will make your releases more professional, reliable, and easier to maintain.