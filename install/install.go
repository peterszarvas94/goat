package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/peterszarvas94/goat/utils"
)

func main() {
	err := utils.ZipFolder("cmd/embed", "cmd/embed.zip", []string{"node_modules", "go.work", "go.work.sum"})
	if err != nil {
		fmt.Println("Error zipping:", err.Error())
		os.Exit(1)
	}

	cmd := exec.Command("go", "install", "-mod", "vendor", "main.go")
	cmd.Dir = "."
	err = cmd.Run()
	if err != nil {
		fmt.Println("Error installing:", err.Error())
		os.Exit(1)
	}

	fmt.Println("Zipping and installation complete!")
}
