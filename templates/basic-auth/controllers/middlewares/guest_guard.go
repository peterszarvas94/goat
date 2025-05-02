package middlewares

import (
	"errors"
	"net/http"
	"scaffhold/controllers/helpers"

	"github.com/peterszarvas94/goat/request"
)

func GuestGuard(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _, err := helpers.CheckAuthStatus(r)
		if err == nil {
			request.ServerError(w, r, errors.New("User should not be logged in"))
			return
		}
		next(w, r)
	}
}
