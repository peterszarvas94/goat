package main

import (
	"fmt"
	"os"

	"github.com/peterszarvas94/goat/utils"
)

func main() {
	err := utils.ZipFolder("cmd/embed", "cmd/embed.zip", []string{"node_modules", "go.work", "go.work.sum", "go.mod", "go.sum", ".gitignore"})
	if err != nil {
		fmt.Println("Error zipping:", err.Error())
		os.Exit(1)
	}

	err = utils.Cmd("go", "install", "-mod=vendor", "./...")
	if err != nil {
		fmt.Println("Error installing:", err.Error())
		os.Exit(1)
	}

	fmt.Println("Zipping and installation complete!")
}
