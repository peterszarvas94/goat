package main

import (
	"fmt"
	"os"

	"github.com/peterszarvas94/goat/constants"
	"github.com/peterszarvas94/goat/utils"
)

func main() {
	err := utils.ZipFolder("cmd/embed", "cmd/embed.zip", []string{"node_modules", "go.work", "go.work.sum", "go.mod", "go.sum", ".gitignore"})
	if err != nil {
		fmt.Println("Error zipping:", err.Error())
		os.Exit(1)
	}

	version := constants.Version

	// Check for uncommitted changes
	if err := utils.Cmd("git", "diff", "--quiet"); err != nil {
		fmt.Println("Uncommitted changes found. Commit before tagging.")
		os.Exit(1)
	}

	// Check if the tag already exists
	err = utils.Cmd("git", "rev-parse", "--verify", fmt.Sprintf("refs/tags/%s", version))
	if err == nil {
		fmt.Printf("Tag %s already exists\n", version)
		os.Exit(1)
	}

	// Create the new tag
	if err := utils.Cmd("git", "tag", version, "-m", version); err != nil {
		fmt.Printf("git tag %s failed: %v\n", version, err)
		os.Exit(1)
	}

	// Push the changes
	if err := utils.Cmd("git", "push"); err != nil {
		fmt.Printf("git push failed: %v\n", err)
		os.Exit(1)
	}

	// Push tags
	if err := utils.Cmd("git", "push", "--tags"); err != nil {
		fmt.Printf("git push --tags failed: %v\n", err)
		os.Exit(1)
	}

	// Success message
	fmt.Printf("Version %s published successfully\n", version)
}
