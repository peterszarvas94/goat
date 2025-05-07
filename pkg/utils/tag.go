package utils

import "strings"

func RemoteTagExists(tag string) (bool, error) {
	output, err := CmdWithOutput("git", "ls-remote", "--tags", "origin")
	if err != nil {
		return false, err
	}
	return strings.Contains(output, "refs/tags/"+tag), nil
}
