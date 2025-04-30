package middlewares

import (
	"context"
	"net/http"
	"scaffhold/db/models"
	"time"

	"github.com/peterszarvas94/goat/csrf"
	"github.com/peterszarvas94/goat/database"
	"github.com/peterszarvas94/goat/logger"
	"github.com/peterszarvas94/goat/request"
)

func AddAuthState(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		db, err := database.Get()
		if err != nil {
			request.ServerError(w, r, err)
			return
		}

		cookie, err := r.Cookie("sessionToken")
		if err != nil {

			// no cookie
			logger.Debug("Cookie not found")
			next(w, r)
			return
		}

		logger.Debug("Cookie found")

		queries := models.New(db)
		session, err := queries.GetSessionByID(context.Background(), cookie.Value)
		if err != nil {
			request.ResetCookie(&w, "sessionToken")

			csrf.Delete(cookie.Value)

			// cookie, but no session -> next
			logger.Debug("Cookie is found, but no session")
			return
		}

		if session.ValidUntil.Before(time.Now()) {
			logger.Debug("Session is expired", "session_id", session.ID)

			err = queries.DeleteSession(context.Background(), session.ID)
			if err != nil {
				request.ServerError(w, r, err)
				return
			}

			csrf.Delete(session.ID)
			next(w, r)
			return
		}

		logger.Debug("Session valid", "session_id", session.ID)

		user, err := queries.GetUserByID(context.Background(), session.UserID)
		if err != nil {
			request.ServerError(w, r, err)
			return
		}

		logger.Debug("User exist", "user_id", user.ID)

		// cookie, session and csrf token, and valid -> next with ctx
		ctx := context.WithValue(r.Context(), "user", &user)
		ctx = context.WithValue(ctx, "session", &session)
		next(w, r.WithContext(ctx))
	}
}
