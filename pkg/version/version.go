package version

import (
	"runtime/debug"
)

// Version is set by GoReleaser during build via ldflags
var Version = "dev"

// Get returns the version, attempting to get it from build info if not set via ldflags
func Get() string {
	if Version != "dev" {
		return Version
	}

	// Try to get version from build info (works with go install github.com/user/repo@version)
	if info, ok := debug.ReadBuildInfo(); ok {
		if info.Main.Version != "(devel)" && info.Main.Version != "" {
			return info.Main.Version
		}
	}

	return "dev"
}
