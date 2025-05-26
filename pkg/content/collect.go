package content

import (
	"bytes"
	"fmt"
	"io/fs"
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	"github.com/peterszarvas94/goat/pkg/constants"
	"github.com/yuin/goldmark"
	emoji "github.com/yuin/goldmark-emoji"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
)

type FileMap map[string]string

// key: file path, value: file content
var Files FileMap

func Setup() (FileMap, error) {
	// clear map
	Files = make(FileMap)

	err := clearHTMLFolder()
	if err != nil {
		return make(FileMap), err
	}

	err = filepath.WalkDir(constants.MarkdownDir, walk())
	if err != nil {
		return make(FileMap), err
	}

	return Files, nil
}

func walk() fs.WalkDirFunc {
	return func(markdownFilePath string, info fs.DirEntry, err error) error {
		if err != nil {
			return fmt.Errorf("failed to access path %s: %w", markdownFilePath, err)
		}

		if info.IsDir() || !strings.HasSuffix(info.Name(), ".md") {
			return nil
		}

		html, err := parseMarkdownFile(markdownFilePath)
		if err != nil {
			return err
		}

		err = writeHtmlFile(markdownFilePath, html)
		if err != nil {
			return err
		}

		return nil
	}
}

func parseMarkdownFile(path string) (string, error) {
	contentBytes, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("failed to read file %s: %w", path, err)
	}

	rawContent := string(contentBytes)

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
		return "", err
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

func writeHtmlFile(filePath, html string) error {
	// content/md/hello/world.md -> hello/world.md
	mdPath := filePath[len(constants.MarkdownDir)+1:]

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
		return err
	}
	defer file.Close()

	_, err = file.WriteString(html)
	if err != nil {
		return err
	}

	Files[route] = htmlPath
	slog.Debug("HTML file generated", "route", route, "path", htmlPath)

	return nil
}
