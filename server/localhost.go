package server

import "strings"

func NewLocalHostUrl(port string) string {
	return strings.Join([]string{"localhost", port}, ":")
}
