package procedures

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/peterszarvas94/goat/logger"
	"github.com/peterszarvas94/goat/request"
)

var count = 0

func GetCount(w http.ResponseWriter, r *http.Request) {
	reqID, ok := r.Context().Value("req_id").(string)
	if reqID == "" || !ok {
		request.ServerError(w, r, errors.New("Request ID is missing"))
		return
	}

	logger.Debug("Count", "req_id", reqID)
	w.Header().Set("Content-Type", "text/html")
	w.Write(fmt.Appendf([]byte{}, "%d", count))
}
