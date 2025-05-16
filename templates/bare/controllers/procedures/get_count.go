package procedures

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/peterszarvas94/goat/pkg/request"
)

var count = 0

func GetCount(w http.ResponseWriter, r *http.Request) {
	reqID, ok := r.Context().Value("req_id").(string)
	if reqID == "" || !ok {
		request.ServerError(w, r, errors.New("Request ID is missing"))
		return
	}

	slog.Debug("Count", "req_id", reqID)
	w.Header().Set("Content-Type", "text/html")
	w.Write(fmt.Appendf([]byte{}, "%d", count))
}
