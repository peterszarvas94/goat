package request

import (
	"fmt"
	"log/slog"
	"net/http"
)

func HttpRedirect(w http.ResponseWriter, r *http.Request, path string, args ...any) {
	slog.Debug(fmt.Sprintf("Redirecting to %s", path), args...)
	http.Redirect(w, r, path, http.StatusMovedPermanently)
}
