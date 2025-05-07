package utils

import (
	"os"
	"strings"
)

func GetVersion() (string, error) {
	file, err := os.ReadFile("VERSION")
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(file)), nil
}
