package uuid

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
)

func New(prefix string) string {
	raw := uuid.New().String()
	return fmt.Sprintf("%s_%s", prefix, strings.ReplaceAll(raw, "-", ""))
}
