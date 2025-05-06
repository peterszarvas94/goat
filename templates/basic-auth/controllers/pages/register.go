package pages

import (
	"errors"
	"net/http"
	"scaffhold/views/pages"

	"github.com/peterszarvas94/goat/pkg/logger"
	"github.com/peterszarvas94/goat/pkg/request"
	"github.com/peterszarvas94/goat/pkg/server"
)

func Register(w http.ResponseWriter, r *http.Request) {
	reqID, ok := r.Context().Value("req_id").(string)
	if reqID == "" || !ok {
		request.ServerError(w, r, errors.New("Request ID is missing"))
		return
	}

	logger.Debug("Rendering register page", "req_id", reqID)
	server.Render(w, r, pages.Register(), http.StatusOK)
}
