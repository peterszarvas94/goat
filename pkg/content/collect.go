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
	"github.com/yuin/goldmark/text"
	"go.abhg.dev/goldmark/frontmatter"
)

type File struct {
	HtmlPath    string
	Content     string
	Frontmatter map[string]any
}

// key is the "route"
type FileMap map[string]File

// key: file path, value: file content
var Files FileMap

var md goldmark.Markdown

func Setup() (*FileMap, error) {
	// clear map
	Files = make(FileMap)

	err := clearHtmlFolder()
	if err != nil {
		return nil, err
	}

	err = filepath.WalkDir(constants.MarkdownDir, walk())
	if err != nil {
		return nil, err
	}

	return &Files, nil
}

func walk() fs.WalkDirFunc {
	return func(markdownFilePath string, info fs.DirEntry, err error) error {
		if err != nil {
			return fmt.Errorf("failed to access path %s: %w", markdownFilePath, err)
		}

		if info.IsDir() || !strings.HasSuffix(info.Name(), ".md") {
			return nil
		}

		markdownContent, err := readFile(markdownFilePath)
		if err != nil {
			return fmt.Errorf("Can not read file %s: %v\n", markdownFilePath, err)
		}

		root := md.Parser().Parse(text.NewReader([]byte(markdownContent)))
		doc := root.OwnerDocument()
		matter := doc.Meta()

		htmlContent, err := mdToHtml(markdownContent)
		if err != nil {
			return fmt.Errorf("Can not parse markdown: %v\n", err)
		}

		route, htmlFilePath := getHtmlInfo(markdownFilePath)

		err = writeHtmlFile(htmlFilePath, htmlContent)
		if err != nil {
			return fmt.Errorf("Can not write html file: %v\n", err)
		}

		Files[route] = File{
			HtmlPath:    htmlFilePath,
			Content:     htmlContent,
			Frontmatter: matter,
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

func mdToHtml(rawContent string) (string, error) {
	var buf bytes.Buffer
	if err := md.Convert([]byte(rawContent), &buf); err != nil {
		return "", fmt.Errorf("failed to convert file: %v", err)
	}

	return buf.String(), nil
}

func clearHtmlFolder() error {
	err := os.RemoveAll(constants.HtmlDir)
	if err != nil {
		return err
	}
	return nil
}

func getHtmlInfo(markdownFilePath string) (route string, htmlPath string) {
	// content/md/hello/world.md -> hello/world.md
	mdPath := markdownFilePath[len(constants.MarkdownDir)+1:]

	// hello/world or hello\world
	barePath := mdPath[:len(mdPath)-len(".md")]

	// hello/world
	route = filepath.ToSlash(barePath)

	// content/html/hello/world.html
	htmlPath = filepath.Join(constants.HtmlDir, barePath+".html")

	return route, htmlPath
}

func writeHtmlFile(htmlFilePath, content string) error {
	// content/html/hello/world.html -> content/html/hello
	htmlFileDir := filepath.Dir(htmlFilePath)

	err := os.MkdirAll(htmlFileDir, 0755)
	if err != nil {
		return err
	}

	file, err := os.Create(htmlFilePath)
	if err != nil {
		return fmt.Errorf("failed to create html file %s: %v", htmlFilePath, err)
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		return fmt.Errorf("failed to write html file %s: %v", htmlFilePath, err)
	}

	slog.Debug("HTML file generated", "path", htmlFilePath)

	return nil
}

func init() {
	md = goldmark.New(
		goldmark.WithExtensions(
			extension.GFM,
			extension.DefinitionList,
			extension.Footnote,
			emoji.Emoji,
			&frontmatter.Extender{
				Mode: frontmatter.SetMetadata,
			},
		),
		goldmark.WithParserOptions(
			parser.WithAttribute(),
		),
	)
}
