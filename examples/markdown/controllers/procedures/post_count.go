package procedures

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/peterszarvas94/goat/pkg/request"
)

func PostCountHandler(w http.ResponseWriter, r *http.Request) {
	reqID, ok := r.Context().Value("req_id").(string)
	if reqID == "" || !ok {
		request.ServerError(w, r, errors.New("Request ID is missing"))
		return
	}

	count++

	slog.Debug("Count increased", "req_id", reqID)
	GetCountHandler(w, r)
}
