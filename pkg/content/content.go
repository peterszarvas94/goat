package content

import (
	"bytes"
	"context"
	"fmt"
	"io/fs"
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	"github.com/a-h/templ"
	"github.com/peterszarvas94/goat/pkg/constants"
	"github.com/yosssi/gohtml"
	"github.com/yuin/goldmark"
	emoji "github.com/yuin/goldmark-emoji"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
	"go.abhg.dev/goldmark/frontmatter"
)

type File struct {
	MarkdownPath string
	HtmlPath     string
	Content      string
	Frontmatter  map[string]any
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

		htmlContent, err := mdToHtml(markdownContent, matter)
		if err != nil {
			return fmt.Errorf("Can not parse markdown: %v\n", err)
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

func mdToHtml(rawContent string, matter map[string]any) (string, error) {
	var htmlContent string

	// no template set, just parse to html
	var buf bytes.Buffer
	if err := md.Convert([]byte(rawContent), &buf); err != nil {
		return "", fmt.Errorf("failed to convert file: %v", err)
	}

	htmlContent = buf.String()

	// template is set, rendering
	if TemplateFunc != nil {
		var buf bytes.Buffer
		TemplateFunc(htmlContent, matter).Render(context.Background(), &buf)
		htmlContent = buf.String()
	}

	if formatting {
		htmlContent = gohtml.Format(htmlContent)
	}

	return htmlContent, nil
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
