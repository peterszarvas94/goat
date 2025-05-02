package request

import (
	"fmt"
	"net/http"

	"github.com/peterszarvas94/goat/pkg/logger"
)

func HxRedirect(w http.ResponseWriter, r *http.Request, path string, args ...any) {
	logger.Debug(fmt.Sprintf("Redirecting to %s", path), args...)
	w.Header().Set("HX-Redirect", path)
	w.WriteHeader(http.StatusMovedPermanently)
	w.Write([]byte{})
}
