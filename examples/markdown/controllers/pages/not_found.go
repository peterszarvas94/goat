package pages

import (
	"errors"
	"log/slog"
	"markdown/views/pages"
	"net/http"

	"github.com/peterszarvas94/goat/pkg/request"
	"github.com/peterszarvas94/goat/pkg/server"
)

func NotFoundPageHandler(w http.ResponseWriter, r *http.Request) {
	reqID, ok := r.Context().Value("req_id").(string)
	if reqID == "" || !ok {
		request.ServerError(w, r, errors.New("Request ID is missing"))
		return
	}

	slog.Debug("Rendering notfound page", "path", r.URL.Path, "req_id", reqID)
	server.Render(w, r, pages.NotFoundPageTemplate(), http.StatusNotFound)
}
