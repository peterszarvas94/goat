package utils

import (
	"strings"
)

func HasUncomittedChanges() (bool, error) {
	output, err := CmdWithOutput("git", "status", "--porcelain")
	if err != nil {
		return false, err
	}
	return strings.TrimSpace(output) != "", nil
}
