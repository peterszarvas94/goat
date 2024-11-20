package cmd

import (
	"fmt"
	"os/exec"
)

// used for local development
func linkLocal(path string) error {
	cmd := exec.Command("go", "mod", "edit", "-replace", fmt.Sprintf("github.com/peterszarvas94/goat=%s", path))
	output, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}

	fmt.Println(string(output))

	cmd = exec.Command("go", "mod", "edit", "-require=github.com/peterszarvas94/goat@v0.0.0-00010101000000-000000000000")

	output, err = cmd.CombinedOutput()
	if err != nil {
		return err
	}

	fmt.Println(string(output))

	return nil
}
