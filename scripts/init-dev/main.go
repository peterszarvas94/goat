package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/peterszarvas94/goat/pkg/dependencies"
	"github.com/peterszarvas94/goat/pkg/utils"
)

func main() {
	original, err := os.Getwd()
	if err != nil {
		fmt.Println("Error changing dir:", err.Error())
		os.Exit(1)
	}

	for template := range dependencies.Templates {
		err := os.Chdir(filepath.Join(original, "cmd", "embed", template))
		if err != nil {
			fmt.Println("Error changing dir:", err.Error())
			os.Exit(1)
		}

		err = utils.Cmd("go", "mod", "init", template)
		if err != nil {
			fmt.Println("Error initializing:", err.Error())
			os.Exit(1)
		}

		fmt.Println("Go initialized")

		dependencies.InstallAll(template)

		err = utils.Cmd("go", "mod", "tidy", template)
		if err != nil {
			fmt.Println("Error initializing:", err.Error())
			os.Exit(1)
		}

		err = utils.Cmd("go", "mod", "vendor", template)
		if err != nil {
			fmt.Println("Error initializing:", err.Error())
			os.Exit(1)
		}

		err = os.Chdir(original)
		if err != nil {
			fmt.Println("Error changing dir:", err.Error())
			os.Exit(1)
		}
	}
}
