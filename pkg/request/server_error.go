package request

import (
	"fmt"
	"log/slog"
	"net/http"
)

func ServerError(w http.ResponseWriter, r *http.Request, err error, args ...any) {
	slog.Error(err.Error(), args...)
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprintln(w, "Internal server error")
}
