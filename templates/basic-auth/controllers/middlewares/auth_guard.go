package middlewares

import (
	"basic-auth/controllers/helpers"
	"net/http"

	"github.com/peterszarvas94/goat/pkg/request"
)

func AuthGuard(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _, err := helpers.CheckAuthStatus(r)
		if err != nil {
			if r.Method == "GET" {
				request.HttpRedirect(w, r, "/login")
				return
			}

			request.Unauthorized(w, r, err)
			return
		}

		next(w, r)
	}
}
