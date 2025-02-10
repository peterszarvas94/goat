package request

import (
	"fmt"
	"net/http"

	"github.com/peterszarvas94/goat/logger"
)

func Conflict(w http.ResponseWriter, r *http.Request, err error, args ...any) {
	logger.Error(err.Error(), args...)
	w.WriteHeader(http.StatusConflict)
	fmt.Fprintln(w, err.Error())
}
