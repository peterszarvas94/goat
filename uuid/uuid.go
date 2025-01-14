package uuid

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/peterszarvas94/goat/assert"
)

func New(prefix string) string {
	assert.Len(3, prefix)
	raw := uuid.New().String()
	return fmt.Sprintf("%s_%s", prefix, strings.ReplaceAll(raw, "-", ""))
}
