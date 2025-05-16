package request

import (
	"fmt"
	"log/slog"
	"net/http"
)

func BadRequest(w http.ResponseWriter, r *http.Request, err error, args ...any) {
	slog.Warn(err.Error(), args...)
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprintln(w, err.Error())
}
