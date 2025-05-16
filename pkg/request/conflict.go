package request

import (
	"fmt"
	"log/slog"
	"net/http"
)

func Conflict(w http.ResponseWriter, r *http.Request, err error, args ...any) {
	slog.Error(err.Error(), args...)
	w.WriteHeader(http.StatusConflict)
	fmt.Fprintln(w, err.Error())
}
