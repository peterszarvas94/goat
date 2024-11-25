package cmd

import (
	"fmt"
	"os/exec"
)

func cmd(name string, args ...string) (string, error) {
	cmd := exec.Command(name, args...)
	fmt.Println(cmd.String())
	output, err := cmd.CombinedOutput()

	return string(output), err
}
