package helpers

import (
	"fmt"
	"strings"

	"github.com/peterszarvas94/goat/config"
)

func GetFileNameFromGooseOutput(output string) (string, error) {
	arr := strings.Split(output, " ")
	if len(arr) < 6 {
		return "", fmt.Errorf("goose output is malformed")
	}

	filename := arr[5]

	if !strings.HasPrefix(filename, config.MigrationsPath) {
		return "", fmt.Errorf("goose output is malformed")
	}

	return strings.TrimSuffix(filename, "\n"), nil
}
