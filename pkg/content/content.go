package content

import (
	"bytes"
	"context"
	"fmt"
	"io/fs"
	"log/slog"
	"maps"
	"os"
	"path/filepath"
	"strings"

	"github.com/a-h/templ"
	"github.com/peterszarvas94/goat/pkg/constants"
	"github.com/peterszarvas94/gohtml"
	"github.com/yuin/goldmark"
	emoji "github.com/yuin/goldmark-emoji"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
	"go.abhg.dev/goldmark/frontmatter"
)

type File struct {
	MarkdownPath  string
	HtmlPath      string
	MdContent     string // The content of the md file
	ParsedContent string // Parsed md to html
	Content       string // Fully rendered html, with the template included, formatted
	Frontmatter   map[string]any
}

// key is the "route"
type FileMap map[string]File

// key: file path, value: file content
var Files FileMap

var md goldmark.Markdown

var TemplateFunc func(htmlContent string, matter map[string]any) templ.Component

func RegisterTemplate(template func(htmlContent string, matter map[string]any) templ.Component) {
	TemplateFunc = template
}

func GetNotFoundFile() (error, *File) {
	for _, file := range maps.All(Files) {
		if file.HtmlPath == constants.NotFoundFile {
			return nil, &file
		}
	}
	return fmt.Errorf("Add a '%s' or '%s'\n", constants.NotFoundTemplate1, constants.NotFoundTemplate2), nil
}

var formatting = true

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

		parsedContent, err := parseMd(markdownContent)
		if err != nil {
			return fmt.Errorf("Can not parse markdown: %v\n", err)
		}

		htmlContent, err := applyTemplate(parsedContent, matter)
		if err != nil {
			return fmt.Errorf("Can not apply template: %v\n", err)
		}

		route, htmlFilePath := getHtmlInfo(markdownFilePath)

		if file, exists := Files[route]; exists {
			return fmt.Errorf("Can not covert file %s, route %s already used for %s\n", markdownFilePath, route, file.MarkdownPath)
		}

		err = writeHtmlFile(htmlFilePath, htmlContent)
		if err != nil {
			return fmt.Errorf("Can not write html file: %v\n", err)
		}

		Files[route] = File{
			MarkdownPath: markdownFilePath,
			HtmlPath:     htmlFilePath,
			MdContent:    markdownContent,
			Content:      htmlContent,
			Frontmatter:  matter,
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

func parseMd(markdownContent string) (string, error) {
	var htmlContent string

	var buf bytes.Buffer
	err := md.Convert([]byte(markdownContent), &buf)
	if err != nil {
		return "", fmt.Errorf("failed to convert file: %v", err)
	}

	htmlContent = buf.String()
	return htmlContent, nil
}

func applyTemplate(parsedContent string, matter map[string]any) (string, error) {
	content := parsedContent

	if TemplateFunc != nil {
		var buf bytes.Buffer
		TemplateFunc(content, matter).Render(context.Background(), &buf)
		content = buf.String()
	}

	if formatting {
		content = gohtml.Format(content)
	}

	return content, nil
}

func clearHtmlFolder() error {
	err := os.RemoveAll(constants.HtmlDir)
	if err != nil {
		return err
	}
	return nil
}

func getHtmlInfo(markdownFilePath string) (route string, htmlPath string) {
	htmlPath = strings.TrimSuffix(markdownFilePath, ".md")
	htmlPath = strings.TrimSuffix(htmlPath, "/index")
	htmlPath = strings.Replace(htmlPath, constants.MarkdownDir, constants.HtmlDir, 1)

	route = strings.TrimPrefix(htmlPath, constants.HtmlDir)
	route = filepath.ToSlash(route)
	if route == "" {
		route = "/"
	}

	htmlPath = fmt.Sprintf("%s/index.html", htmlPath)

	return route, htmlPath
}

func writeHtmlFile(htmlFilePath, content string) error {
	htmlFileDir := filepath.Dir(htmlFilePath)

	err := os.MkdirAll(htmlFileDir, 0755)
	if err != nil {
		return err
	}

	if _, err := os.Stat(htmlFilePath); err == nil {
		return fmt.Errorf("html file %s already exists", htmlFilePath)
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

func SetIndentString(indent string) {
	gohtml.IndentString = indent
}

func SetCondense(condense bool) {
	gohtml.Condense = condense
}

func SetFormatting(format bool) {
	formatting = format
}

func init() {
	gohtml.Condense = true
	gohtml.IndentString = "\t"

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
