package utils

import (
	"fmt"
	"strings"

	"github.com/peterszarvas94/goat/constants"
)

func getFileNameFromGooseOutput(output string) (string, error) {
	arr := strings.Split(output, " ")
	if len(arr) < 6 {
		return "", fmt.Errorf("goose output is malformed")
	}

	filename := arr[5]

	if !strings.HasPrefix(filename, constants.MigrationsDir) {
		return "", fmt.Errorf("goose output is malformed")
	}

	return strings.TrimSuffix(filename, "\n"), nil
}
