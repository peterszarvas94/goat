package utils

import (
	"os"
	"strings"
)

func RemoveReplaceLines(goModPath string) error {
	data, err := os.ReadFile(goModPath)
	if err != nil {
		return err
	}

	var kept []string
	for line := range strings.SplitSeq(string(data), "\n") {
		if !strings.HasPrefix(strings.TrimSpace(line), "replace") {
			kept = append(kept, line)
		}
	}

	return os.WriteFile(goModPath, []byte(strings.Join(kept, "\n")), 0644)
}
