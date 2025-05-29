package content

import (
	"bytes"
	"fmt"
	"io/fs"
	"log/slog"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/adrg/frontmatter"
	"github.com/peterszarvas94/goat/pkg/constants"
	"github.com/yuin/goldmark"
	emoji "github.com/yuin/goldmark-emoji"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
)

type File struct {
	HtmlPath    string
	Content     string
	Frontmatter any
}

// key is the "route"
type FileMap map[string]File

// key: file path, value: file content
var Files FileMap

func Setup(matter any) (FileMap, error) {
	// clear map
	Files = make(FileMap)

	err := clearHTMLFolder()
	if err != nil {
		return make(FileMap), err
	}

	err = filepath.WalkDir(constants.MarkdownDir, walk(matter))
	if err != nil {
		return make(FileMap), err
	}

	return Files, nil
}

func walk(matter any) fs.WalkDirFunc {
	return func(markdownFilePath string, info fs.DirEntry, err error) error {
		if err != nil {
			return fmt.Errorf("failed to access path %s: %w", markdownFilePath, err)
		}

		if info.IsDir() || !strings.HasSuffix(info.Name(), ".md") {
			return nil
		}

		content, err := readFile(markdownFilePath)
		if err != nil {
			return fmt.Errorf("Can not read file %s: %v\n", markdownFilePath, err)
		}

		var matter2 any

		if reflect.TypeOf(matter).Kind() == reflect.Ptr &&
			reflect.TypeOf(matter).Elem().Kind() == reflect.Struct {

			originalValue := reflect.ValueOf(matter).Elem()
			newValue := reflect.New(originalValue.Type())
			newValue.Elem().Set(originalValue)

			matter2 = newValue.Interface()
		}

		body, err := frontmatter.MustParse(strings.NewReader(content), matter2)
		if err != nil {
			return fmt.Errorf("Can not parse frontmatter: %v\n", err)
		}
		content = string(body)

		htmlContent, err := parseMarkdown(content)
		if err != nil {
			return fmt.Errorf("Can not parse markdown: %v\n", err)
		}

		err = writeHtmlFile(markdownFilePath, htmlContent, matter2)
		if err != nil {
			return fmt.Errorf("Can not write html file: %v\n", err)
		}

		return nil
	}
}

func readFile(path string) (string, error) {
	contentBytes, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("failed to read file %s: %w", path, err)
	}

	return string(contentBytes), nil
}

func parseMarkdown(rawContent string) (string, error) {
	markdown := goldmark.New(
		goldmark.WithExtensions(
			extension.GFM,
			extension.DefinitionList,
			extension.Footnote,
			emoji.Emoji,
		),
		goldmark.WithParserOptions(
			parser.WithAttribute(),
		),
	)

	var buf bytes.Buffer
	if err := markdown.Convert([]byte(rawContent), &buf); err != nil {
		return "", fmt.Errorf("failed to convert file: %v", err)
	}

	return buf.String(), nil
}

func clearHTMLFolder() error {
	err := os.RemoveAll(constants.HtmlDir)
	if err != nil {
		return err
	}
	return nil
}

func writeHtmlFile(markdownFilePath, content string, matter any) error {
	// content/md/hello/world.md -> hello/world.md
	mdPath := markdownFilePath[len(constants.MarkdownDir)+1:]

	// hello/world or hello\world
	barePath := mdPath[:len(mdPath)-len(".md")]

	// Normalize the slashes to forward slashes for route
	route := filepath.ToSlash(barePath)

	// content/html/hello/world.html
	htmlPath := filepath.Join(constants.HtmlDir, barePath+".html")

	// content/html/hello/world.html -> content/html/hello
	htmlFileDir := filepath.Dir(htmlPath)

	err := os.MkdirAll(htmlFileDir, 0755)
	if err != nil {
		return err
	}

	file, err := os.Create(htmlPath)
	if err != nil {
		return fmt.Errorf("failed to create html file %s: %v", htmlPath, err)
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		return fmt.Errorf("failed to write html file %s: %v", htmlPath, err)
	}

	Files[route] = File{
		HtmlPath:    htmlPath,
		Content:     content,
		Frontmatter: matter,
	}

	slog.Debug("HTML file generated", "route", route, "path", htmlPath)

	return nil
}
