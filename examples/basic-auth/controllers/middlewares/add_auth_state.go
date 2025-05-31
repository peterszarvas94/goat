package middlewares

import (
	"basic-auth/db/models"
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/peterszarvas94/goat/pkg/csrf"
	"github.com/peterszarvas94/goat/pkg/database"
	"github.com/peterszarvas94/goat/pkg/request"
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
			slog.Debug("Cookie not found")
			next(w, r)
			return
		}

		slog.Debug("Cookie found")

		queries := models.New(db)
		session, err := queries.GetSessionByID(context.Background(), cookie.Value)
		if err != nil {
			request.ResetCookie(&w, "sessionToken")

			csrf.Delete(cookie.Value)

			// cookie, but no session -> next
			slog.Debug("Cookie is found, but no session")
			return
		}

		if session.ValidUntil.Before(time.Now()) {
			slog.Debug("Session is expired", "session_id", session.ID)

			err = queries.DeleteSession(context.Background(), session.ID)
			if err != nil {
				request.ServerError(w, r, err)
				return
			}

			csrf.Delete(session.ID)
			next(w, r)
			return
		}

		slog.Debug("Session valid", "session_id", session.ID)

		user, err := queries.GetUserByID(context.Background(), session.UserID)
		if err != nil {
			request.ServerError(w, r, err)
			return
		}

		slog.Debug("User exist", "user_id", user.ID)

		// cookie, session and csrf token, and valid -> next with ctx
		ctx := context.WithValue(r.Context(), "user", &user)
		ctx = context.WithValue(ctx, "session", &session)
		next(w, r.WithContext(ctx))
	}
}
