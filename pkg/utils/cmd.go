package utils

import (
	"os"
	"os/exec"
)

func Cmd(base string, args ...string) error {
	cmd := exec.Command(base, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func CmdWithOutput(base string, args ...string) (string, error) {
	cmd := exec.Command(base, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	output, err := cmd.Output()
	return string(output), err
}
