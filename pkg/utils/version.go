package utils

import "github.com/peterszarvas94/goat/pkg/version"

func GetVersion() (string, error) {
	return version.Version, nil
}
