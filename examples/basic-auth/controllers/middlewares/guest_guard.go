package middlewares

import (
	"basic-auth/controllers/helpers"
	"errors"
	"fmt"
	"net/http"

	"github.com/peterszarvas94/goat/pkg/request"
)

func GuestGuard(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _, err := helpers.CheckAuthStatus(r)
		if err == nil {
			if r.Method == "GET" {
				fmt.Println("here")
				request.HttpRedirect(w, r, "/")
				return
			}

			request.ServerError(w, r, errors.New("User should not be logged in"))
			return
		}
		next(w, r)
	}
}
