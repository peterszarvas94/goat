package procedures

import (
	"errors"
	"net/http"

	"github.com/peterszarvas94/goat/pkg/logger"
	"github.com/peterszarvas94/goat/pkg/request"
)

func PostCount(w http.ResponseWriter, r *http.Request) {
	reqID, ok := r.Context().Value("req_id").(string)
	if reqID == "" || !ok {
		request.ServerError(w, r, errors.New("Request ID is missing"))
		return
	}

	count++

	logger.Debug("Count increased", "req_id", reqID)
	GetCount(w, r)
}
