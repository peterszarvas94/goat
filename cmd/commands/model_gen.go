package commands

import (
	"fmt"

	"github.com/peterszarvas94/goat/cmd/helpers"
)

func ModelGen() error {
	output, err := helpers.Cmd("sqlc", "generate")
	fmt.Println(output)
	return err
}
