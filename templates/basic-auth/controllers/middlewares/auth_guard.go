package middlewares

import (
	"net/http"
	"scaffhold/controllers/helpers"

	"github.com/peterszarvas94/goat/request"
)

func AuthGuard(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _, err := helpers.CheckAuthStatus(r)
		if err != nil {
			request.Unauthorized(w, r, err)
			return
		}

		next(w, r)
	}
}
