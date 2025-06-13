package content

import (
	"fmt"
	"testing"

	"github.com/peterszarvas94/goat/pkg/constants"
)

func TestGetHtmlInfo(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		wantRoute string
		wantPath  string
	}{
		{
			name:      "index",
			input:     fmt.Sprintf("%s/index.md", constants.MarkdownDir),
			wantRoute: "/",
			wantPath:  fmt.Sprintf("%s/index.html", constants.HtmlDir),
		},
		{
			name:      "nested_index",
			input:     fmt.Sprintf("%s/hello/world/index.md", constants.MarkdownDir),
			wantRoute: "/hello/world",
			wantPath:  fmt.Sprintf("%s/hello/world/index.html", constants.HtmlDir),
		},
		{
			name:      "named",
			input:     fmt.Sprintf("%s/hello/world.md", constants.MarkdownDir),
			wantRoute: "/hello/world",
			wantPath:  fmt.Sprintf("%s/hello/world/index.html", constants.HtmlDir),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRoute, gotPath := getHtmlInfo(tt.input)

			if tt.wantRoute != gotRoute {
				t.Errorf("Got %s, want %s", gotRoute, tt.wantRoute)
			}

			if tt.wantPath != gotPath {
				t.Errorf("Got %s, want %s", gotPath, tt.wantPath)
			}
		})
	}
}
