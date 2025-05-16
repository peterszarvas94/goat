package request

import (
	"fmt"
	"log/slog"
	"net/http"
)

func HxRedirect(w http.ResponseWriter, r *http.Request, path string, args ...any) {
	slog.Debug(fmt.Sprintf("Redirecting to %s", path), args...)
	w.Header().Set("HX-Redirect", path)
	w.WriteHeader(http.StatusMovedPermanently)
	w.Write([]byte{})
}
