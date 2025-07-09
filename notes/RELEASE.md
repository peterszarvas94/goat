# Release Process with GoReleaser

## Quick Start

### Local Testing
```bash
# Test release configuration
make release-test

# Build snapshot locally  
make release-local
```

### Production Release
```bash
# Create and push tag
git tag v1.2.3
git push origin v1.2.3

# GitHub Actions automatically handles the rest
```

## What Happens

1. **Tag Push**: Triggers GitHub Actions workflow
2. **GoReleaser**: Builds for multiple platforms (Linux, macOS, Windows)
3. **GitHub Release**: Creates release with binaries and changelog
4. **Version Injection**: Binary contains correct version via ldflags

## Files Created

- `.goreleaser.yaml` - GoReleaser configuration
- `.github/workflows/release.yml` - GitHub Actions workflow
- Updated `Makefile` with release commands

## Migration from Old Publish Script

**Before**: `make publish` (manual process)
**After**: `git tag v1.2.3 && git push origin v1.2.3` (automated)

The old `publish/main.go` script is no longer needed for releases.

## Version Management

- Development: `goat version` shows "dev"
- Released binary: `goat version` shows actual version (e.g., "v1.2.3")
- Version injected at build time via ldflags

## Supported Platforms

- Linux (amd64, arm64)
- macOS (amd64, arm64) 
- Windows (amd64, arm64)

All binaries are statically linked (CGO_ENABLED=0) for maximum compatibility.