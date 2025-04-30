package helpers

import (
	"fmt"
	"os/exec"
)

func Cmd(name string, args ...string) (string, error) {
	cmd := exec.Command(name, args...)
	fmt.Println(cmd.String())
	output, err := cmd.CombinedOutput()

	fmt.Println(string(output))

	return string(output), err
}
