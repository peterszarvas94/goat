package middlewares

import (
	"errors"
	"net/http"
	"scaffhold/db/models"

	"github.com/peterszarvas94/goat/csrf"
	"github.com/peterszarvas94/goat/request"
)

func CSRFGuard(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctxSession, ok := r.Context().Value("session").(*models.Session)
		if ctxSession == nil || !ok {
			request.ServerError(w, r, errors.New("Not logged in"))
			return
		}

		err := r.ParseForm()
		if err != nil {
			request.ServerError(w, r, err)
			return
		}

		csrfToken := r.FormValue("csrf_token")

		err = csrf.Validate(ctxSession.ID, csrfToken)
		if err != nil {
			request.ServerError(w, r, errors.New("CSRF token is invalid"))
			return
		}

		next(w, r)
	}
}
