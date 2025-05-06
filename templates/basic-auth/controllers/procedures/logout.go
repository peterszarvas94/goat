package procedures

import (
	"context"
	"errors"
	"net/http"
	"scaffhold/db/models"

	"github.com/peterszarvas94/goat/pkg/csrf"
	"github.com/peterszarvas94/goat/pkg/database"
	"github.com/peterszarvas94/goat/pkg/logger"
	"github.com/peterszarvas94/goat/pkg/request"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	reqID, ok := r.Context().Value("req_id").(string)
	if reqID == "" || !ok {
		request.ServerError(w, r, errors.New("Request ID is missing"))
		return
	}

	cookie, err := r.Cookie("sessionToken")
	if err != nil {
		request.ServerError(w, r, err, "req_id", reqID)
		return
	}

	logger.Debug("Cookie found", "req_id", reqID, "session_id", cookie.Value)

	db, err := database.Get()
	if err != nil {
		request.ServerError(w, r, err, "req_id", reqID)
		return
	}

	queries := models.New(db)
	err = queries.DeleteSession(context.Background(), cookie.Value)
	if err != nil {
		request.ServerError(w, r, err, "req_id", reqID)
		return
	}

	request.ResetCookie(&w, "sessionToken")

	csrf.Delete(cookie.Value)

	logger.Debug("Logged out", "req_id", reqID)
	request.HxRedirect(w, r, "/", "req_id", reqID)
}
