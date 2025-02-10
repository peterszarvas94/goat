package request

import (
	"fmt"
	"net/http"

	"github.com/peterszarvas94/goat/logger"
)

func BadRequest(w http.ResponseWriter, r *http.Request, err error, args ...any) {
	logger.Warn(err.Error(), args...)
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprintln(w, err.Error())
}
