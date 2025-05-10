package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/peterszarvas94/goat/pkg/utils"
)

func main() {
	// checking version
	version, err := utils.GetVersion()
	if err != nil {
		fmt.Printf("Can not get version: %s", err.Error())
		os.Exit(1)
	}

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

	// changing go.mod files for templates:
	// - backup to go.mod.bkp
	// - replacing goat version
	// - removing "replace" directives
	subfolders, err := utils.GetSubfolders("templates")
	for _, folder := range subfolders {
		modFilePath := filepath.Join(folder, "go.mod")
		backupModFilePath := filepath.Join(folder, "go.mod.bkp")

		err := utils.CopyFile(modFilePath, backupModFilePath)
		if err != nil {
			fmt.Printf("Can not backup modfile: %s\n", folder)
			os.Exit(1)
		}

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
			if strings.Contains(line, "github.com/peterszarvas94/goat/pkg v") {
				start := strings.Index(line, "github.com/peterszarvas94/goat/pkg")
				if start != -1 {
					parts := strings.Fields(line[start:])
					if len(parts) == 2 {
						line = line[:start] + parts[0] + " " + version
					}
				}
			}

			if strings.Contains(line, "replace") && !strings.HasPrefix(line, "// ") {
				newContent.WriteString("// ")
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

		fmt.Printf("Version replaced for folder %s to %s\n", folder, version)
	}

	// git stuff
	err = utils.Cmd("git", "add", ".")
	if err != nil {
		fmt.Printf("Error with \"git add .\": %s\n", err.Error())
		os.Exit(1)
	}

	err = utils.Cmd("git", "commit", "-m", fmt.Sprintf("publish: %s", version))
	if err != nil {
		fmt.Printf("Error with \"git commit\": %s\n", err.Error())
		os.Exit(1)
	}

	err = utils.Cmd("git", "push")
	if err != nil {
		fmt.Printf("Error with \"git push\": %s\n", err.Error())
		os.Exit(1)
	}

	err = utils.Cmd("git", "push", "--tags")
	if err != nil {
		fmt.Printf("Error with \"git push --tags\": %s\n", err.Error())
		os.Exit(1)
	}

	// restore modfiles
	for _, folder := range subfolders {
		modFilePath := filepath.Join(folder, "go.mod")
		backupModFilePath := filepath.Join(folder, "go.mod.bkp")

		err = os.Remove(modFilePath)
		if err != nil {
			fmt.Printf("Can not remove modfile: %s\n", err.Error())
			os.Exit(1)
		}

		err = utils.CopyFile(backupModFilePath, modFilePath)
		if err != nil {
			fmt.Printf("Can not restore modfile: %s\n", err.Error())
			os.Exit(1)
		}

		err = utils.Cmd("go", "mod", "tidy")
		if err != nil {
			fmt.Printf("Error with \"git push --tags\": %s\n", err.Error())
			os.Exit(1)
		}
	}
}
