package content

import (
	"fmt"
	"maps"
)

func GetCategorizedFiles(category string) FileMap {
	var res = make(FileMap)

	for route, file := range maps.All(Files) {
		c, err := GetCategory(file)
		if err != nil {
			fmt.Printf("Could not get category from file '%s':\n%v", file.MarkdownPath, err)
			return make(FileMap)
		}

		if c == category {
			res[route] = file
		}
	}
	return res
}

func GetCategory(file File) (string, error) {
	for key, value := range maps.All(file.Frontmatter) {
		if key != "category" {
			continue
		}

		category, ok := value.(string)
		if !ok {
			return "", fmt.Errorf("The 'category' value in frontmatter is not a string in file '%s'\n", file.MarkdownPath)
		}

		return category, nil
	}
	return "", nil
}

func GetCategorySafe(file File) string {
	category, _ := GetCategory(file)
	return category
}
