package pages

import (
	"basic-auth/views/pages"
	"errors"
	"log/slog"
	"net/http"

	"github.com/peterszarvas94/goat/pkg/request"
	"github.com/peterszarvas94/goat/pkg/server"
)

func Register(w http.ResponseWriter, r *http.Request) {
	reqID, ok := r.Context().Value("req_id").(string)
	if reqID == "" || !ok {
		request.ServerError(w, r, errors.New("Request ID is missing"))
		return
	}

	slog.Debug("Rendering register page", "req_id", reqID)
	server.Render(w, r, pages.Register(), http.StatusOK)
}
