package procedures

import (
	"basic-auth/controllers/helpers"
	"basic-auth/db/models"
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/peterszarvas94/goat/pkg/csrf"
	"github.com/peterszarvas94/goat/pkg/database"
	"github.com/peterszarvas94/goat/pkg/request"
)

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	reqID, ok := r.Context().Value("req_id").(string)
	if reqID == "" || !ok {
		helpers.ServerError(w, r, []string{"Request ID is missing"}, true)
		return
	}

	cookie, err := r.Cookie("sessionToken")
	if err != nil {
		helpers.ServerError(w, r, []string{err.Error()}, true, "req_id", reqID)
		return
	}

	slog.Debug("Cookie found", "req_id", reqID, "session_id", cookie.Value)

	db, err := database.Get()
	if err != nil {
		helpers.ServerError(w, r, []string{fmt.Sprintf("Can not get database: %s", err.Error())}, true, "req_id", reqID)
		return
	}

	queries := models.New(db)
	err = queries.DeleteSession(context.Background(), cookie.Value)
	if err != nil {
		helpers.ServerError(w, r, []string{err.Error()}, true, "req_id", reqID)
		return
	}

	request.ResetCookie(&w, "sessionToken")

	csrf.Delete(cookie.Value)

	slog.Debug("Logged out", "req_id", reqID)
	request.HxRedirect(w, r, "/", "req_id", reqID)
}
