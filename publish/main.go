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
	tag, err := utils.GetVersion()
	if err != nil {
		fmt.Printf("Can not get version: %s", err.Error())
		os.Exit(1)
	}

	fmt.Printf("Publishing version: %s\n", tag)

	ok, err := utils.RemoteTagExists(tag)
	if err != nil {
		fmt.Printf("Error checking tag %s: %s\n", tag, err.Error())
	}
	if !ok {
		fmt.Printf("Tag does not exists yet: %s \n", tag)
	} else {
		fmt.Printf("Tag already exists: %s \n", tag)
		os.Exit(1)
	}

	subfolders, err := utils.GetSubfolders("templates")
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
			if strings.Contains(line, "github.com/peterszarvas94/goat/pkg v") {
				start := strings.Index(line, "github.com/peterszarvas94/goat/pkg")
				if start != -1 {
					parts := strings.Fields(line[start:])
					if len(parts) == 2 {
						line = line[:start] + parts[0] + " " + tag
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

		fmt.Printf("Version replaced for folder %s to %s\n", folder, tag)
	}

	// TODO: git add .
	// TODO: git commit -m "publish: version"
	// TODO: git push
	// TODO: git push --tags
	// TODO: restore original go.mod files

	///////////// OLD CODE: /////////////

	// Parse the folder flag
	// folderFlag := flag.String("folder", "", "Folder to publish (e.g., cli/xy)")
	// flag.Parse()
	//
	// if *folderFlag == "" {
	// 	fmt.Println("Please specify a folder using --folder")
	// 	os.Exit(1)
	// }

	//
	// folder := *folderFlag
	// fp := filepath.Join(strings.Split(folder, "/")...)
	// goMod := filepath.Join(fp, "go.mod")
	// localGoMod := filepath.Join(fp, "go.mod.local")
	//
	// // Backup go.mod to go.mod.local
	// if err := utils.CopyFile(goMod, localGoMod); err != nil {
	// 	fmt.Printf("Error backing up go.mod: %v\n", err)
	// 	os.Exit(1)
	// }
	//
	// // Remove replace lines from go.mod
	// if err := utils.RemoveReplaceLines(goMod); err != nil {
	// 	fmt.Printf("Error removing replace lines: %v\n", err)
	// 	os.Exit(1)
	// }
	//
	// // Commit the changes
	// if err := utils.Cmd("git", "add", "."); err != nil {
	// 	fmt.Printf("Error running git add: %v\n", err)
	// 	os.Exit(1)
	// }
	// if err := utils.Cmd("git", "commit", "--amend"); err != nil {
	// 	fmt.Printf("Error running git commit --amend: %v\n", err)
	// 	os.Exit(1)
	// }
	//
	// // Tag the new version
	// tag := fmt.Sprintf("%s%s", folder, version)
	// if err := utils.Cmd("git", "tag", tag, "-m", tag); err != nil {
	// 	fmt.Printf("Error creating git tag: %v\n", err)
	// 	os.Exit(1)
	// }
	//
	// // Push the changes and tags
	// if err := utils.Cmd("git", "push"); err != nil {
	// 	fmt.Printf("Error pushing to git: %v\n", err)
	// 	os.Exit(1)
	// }
	// if err := utils.Cmd("git", "push", "--tags"); err != nil {
	// 	fmt.Printf("Error pushing tags to git: %v\n", err)
	// 	os.Exit(1)
	// }
	//
	// // Restore the original go.mod
	// if err := utils.CopyFile(localGoMod, goMod); err != nil {
	// 	fmt.Printf("Error restoring go.mod: %v\n", err)
	// 	os.Exit(1)
	// }
	//
	// // Remove the backup
	// if err := os.Remove(localGoMod); err != nil {
	// 	fmt.Printf("Error removing go.mod.local: %v\n", err)
	// 	os.Exit(1)
	// }
	//
	// fmt.Printf("Version %s published successfully\n", version)
}
