package main

import (
	_ "embed"

	"github.com/peterszarvas94/goat/cli/cmd"
)

//go:embed VERSION
var version string

func main() {
	cmd.Execute(version)
}
