package pages

import (
	"markdown/views/pages"
	"net/http"

	"github.com/peterszarvas94/goat/pkg/content"
	"github.com/peterszarvas94/goat/pkg/server"
)

func CategoryPageHandler(w http.ResponseWriter, r *http.Request) {
	catagory := r.PathValue("category")

	files := content.GetCategorizedFiles(catagory)

	server.Render(w, r, pages.CategoryPageTemplate(catagory, files), http.StatusOK)
}
