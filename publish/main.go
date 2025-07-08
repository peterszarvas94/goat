package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/peterszarvas94/goat/pkg/constants"
	"github.com/peterszarvas94/goat/pkg/utils"
	"github.com/peterszarvas94/goat/pkg/version"
)

func main() {
	version := version.Version

	fmt.Printf("Publishing version: %s\n", version)

	// checking tag
	ok, err := utils.RemoteTagExists(version)
	if err != nil {
		fmt.Printf("Error checking tag %s: %s\n", version, err.Error())
	}
	if !ok {
		fmt.Printf("Tag does not exists yet: %s \n", version)
	} else {
		fmt.Printf("Tag already exists: %s \n", version)
		os.Exit(1)
	}

	// check git status
	has, err := utils.HasUncomittedChanges()
	if err != nil {
		fmt.Printf("Error checking uncommitted changes %s: %s\n", version, err.Error())
	}
	if has {
		fmt.Println("You have uncomitted changes")
		os.Exit(1)
	} else {
		fmt.Println("You dont have uncomitted changes")
	}

	// changing go.mod files for examples:
	// - replacing goat version
	// - removing "replace" directives
	subfolders, err := utils.GetSubfolders(constants.ExamplesDir)
	if err != nil {
		fmt.Printf("Can not get subfolders: %v\n", err)
		os.Exit(1)
	}

	for _, folder := range subfolders {
		modFilePath := filepath.Join(folder, "go.mod")

		modFile, err := os.Open(modFilePath)
		if err != nil {
			fmt.Printf("No modfile found in: %s\n", folder)
			os.Exit(1)
		}
		defer modFile.Close()

		var newContent strings.Builder
		scanner := bufio.NewScanner(modFile)
		for scanner.Scan() {
			line := scanner.Text()
			start := strings.Index(line, "github.com/peterszarvas94/goat")
			if start != -1 {
				parts := strings.Fields(line[start:])
				if len(parts) == 2 {
					line = line[:start] + parts[0] + " " + version
					fmt.Printf("Updated goat version number in file: %s\n", modFilePath)
				}
			}

			if strings.Contains(line, "replace") && !strings.HasPrefix(line, "// ") {
				line = "// " + line
				fmt.Printf("Commended out replace directive in file: %s\n", modFilePath)
			}
			newContent.WriteString(line)
			newContent.WriteString("\n")
		}

		if err := scanner.Err(); err != nil {
			fmt.Printf("Error reading modfile: %s\n", err.Error())
			os.Exit(1)
		}

		err = os.WriteFile(modFilePath, []byte(newContent.String()), 0644)
		if err != nil {
			fmt.Printf("Error writing modfile: %s\n", err.Error())
			os.Exit(1)
		}
	}

	// git stuff
	err = utils.Cmd("git", "add", ".")
	if err != nil {
		fmt.Printf("Error with \"git add .\": %s\n", err.Error())
		os.Exit(1)
	}

	fmt.Println("Staged files")

	lastCommitMsg, err := utils.CmdWithOutput("git", "log", "-1", "--pretty=format:%s")
	if err != nil {
		fmt.Printf("Error with \"git log\": %s\n", err.Error())
		os.Exit(1)
	}

	err = utils.Cmd("git", "commit", "--amend", "-m", fmt.Sprintf("%s\npublish: %s", lastCommitMsg, version))
	if err != nil {
		fmt.Printf("Error with \"git commit\": %s\n", err.Error())
		os.Exit(1)
	}

	fmt.Println("Commited files")

	err = utils.Cmd("git", "push")
	if err != nil {
		fmt.Printf("Error with \"git push\": %s\n", err.Error())
		os.Exit(1)
	}

	fmt.Println("Pushed files")

	err = utils.Cmd("git", "tag", version, "-m", version)
	if err != nil {
		fmt.Printf("Error with \"git tag\": %s\n", err.Error())
		os.Exit(1)
	}

	err = utils.Cmd("git", "push", "--tags")
	if err != nil {
		fmt.Printf("Error with \"git push --tags\": %s\n", err.Error())
		os.Exit(1)
	}

	fmt.Println("Pushed tags")

	// add "replace" to modfiles
	for _, folder := range subfolders {
		modFilePath := filepath.Join(folder, "go.mod")

		modFile, err := os.Open(modFilePath)
		if err != nil {
			fmt.Printf("No modfile found in: %s\n", folder)
			os.Exit(1)
		}
		defer modFile.Close()

		var newContent strings.Builder
		scanner := bufio.NewScanner(modFile)
		for scanner.Scan() {
			line := scanner.Text()
			start := strings.Index(line, "// replace github.com/peterszarvas94/goat")
			if start != -1 {
				line = strings.TrimPrefix(line, "// ")
				fmt.Printf("Replace directive restored in modfile: %s\n", modFilePath)
			}

			newContent.WriteString(line)
			newContent.WriteString("\n")
		}

		if err := scanner.Err(); err != nil {
			fmt.Printf("Error reading mod file: %s\n", err.Error())
			os.Exit(1)
		}

		err = os.WriteFile(modFilePath, []byte(newContent.String()), 0644)
		if err != nil {
			fmt.Printf("Error writing mod file: %s\n", err.Error())
			os.Exit(1)
		}

		err = utils.Cmd("go", "mod", "tidy")
		if err != nil {
			fmt.Printf("Error with \"go mod tidy\": %s\n", err.Error())
			os.Exit(1)
		}

		fmt.Printf("Tidied folder: %s\n", folder)

	}
}
