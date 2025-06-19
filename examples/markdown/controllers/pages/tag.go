package pages

import (
	"markdown/views/pages"
	"net/http"

	"github.com/peterszarvas94/goat/pkg/content"
	"github.com/peterszarvas94/goat/pkg/server"
)

func TagPageHandler(w http.ResponseWriter, r *http.Request) {
	tag := r.PathValue("tag")

	files := content.GetTaggedFiles(tag)

	server.Render(w, r, pages.TagPageTemplate(files), http.StatusOK)
}
