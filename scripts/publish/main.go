package main

import (
	"fmt"
	"os"

	"github.com/peterszarvas94/goat/pkg/constants"
	"github.com/peterszarvas94/goat/pkg/utils"
)

func main() {
	// Check for uncommitted changes
	if err := utils.Cmd("git", "diff", "--quiet"); err != nil {
		fmt.Println("Uncommitted changes found. Commit before tagging.")
		os.Exit(1)
	}

	// Check if the tag already exists
	err := utils.Cmd("git", "rev-parse", "--verify", fmt.Sprintf("refs/tags/%s", constants.Version))
	if err == nil {
		fmt.Printf("Tag %s already exists\n", constants.Version)
		os.Exit(1)
	}

	// Create the new tag
	if err := utils.Cmd("git", "tag", constants.Version, "-m", constants.Version); err != nil {
		fmt.Printf("git tag %s failed: %v\n", constants.Version, err)
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
	fmt.Printf("Version %s published successfully\n", constants.Version)
}
