package pages

import (
	"errors"
	"net/http"
	"scaffhold/views/pages"

	"github.com/peterszarvas94/goat/logger"
	"github.com/peterszarvas94/goat/request"
	"github.com/peterszarvas94/goat/server"
)

func NotFound(w http.ResponseWriter, r *http.Request) {
	reqID, ok := r.Context().Value("req_id").(string)
	if reqID == "" || !ok {
		request.ServerError(w, r, errors.New("Request ID is missing"))
		return
	}

	logger.Debug("Rendering notfound page", "path", r.URL.Path, "req_id", reqID)
	server.Render(w, r, pages.NotFound(), http.StatusNotFound)
}
