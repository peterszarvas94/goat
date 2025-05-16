package request

import (
	"fmt"
	"log/slog"
	"net/http"
)

func Unauthorized(w http.ResponseWriter, r *http.Request, err error, args ...any) {
	slog.Error(err.Error(), args...)
	w.WriteHeader(http.StatusUnauthorized)
	fmt.Fprintln(w, "Unauthorized")
}
