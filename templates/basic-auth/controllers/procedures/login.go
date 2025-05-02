package procedures

import (
	"context"
	"errors"
	"net/http"
	"time"

	"scaffhold/config"
	"scaffhold/db/models"

	"github.com/peterszarvas94/goat/csrf"
	"github.com/peterszarvas94/goat/database"
	"github.com/peterszarvas94/goat/hash"
	"github.com/peterszarvas94/goat/logger"
	"github.com/peterszarvas94/goat/request"
	"github.com/peterszarvas94/goat/uuid"
)

func Login(w http.ResponseWriter, r *http.Request) {
	reqID, ok := r.Context().Value("req_id").(string)
	if reqID == "" || !ok {
		request.ServerError(w, r, errors.New("Request ID is missing"))
		return
	}

	if err := r.ParseForm(); err != nil {
		request.ServerError(w, r, err, "req_id", reqID)
		return
	}

	email := r.FormValue("email")
	if email == "" {
		request.BadRequest(w, r, errors.New("Email can not be empty"), "req_id", reqID)
		return
	}

	password := r.FormValue("password")
	if password == "" {
		request.BadRequest(w, r, errors.New("Password can not be empty"), "req_id", reqID)
		return
	}

	db, err := database.Get()
	if err != nil {
		request.ServerError(w, r, err, "req_id", reqID)
		return
	}

	queries := models.New(db)
	user, err := queries.GetUserByEmail(context.Background(), email)
	if err != nil {
		request.Unauthorized(w, r, errors.New("User with this email not found"), "req_id", reqID)
		return
	}

	valid := hash.VerifyPassword(password, user.Password)
	if !valid {
		request.Unauthorized(w, r, errors.New("Bad credentials"), "req_id", reqID)
		return
	}

	logger.Debug("Credentials are valid", "req_id", reqID)

	sessionId := uuid.New("ses")
	session, err := queries.CreateSession(context.Background(), models.CreateSessionParams{
		ID:         sessionId,
		UserID:     user.ID,
		ValidUntil: time.Now().Add(24 * time.Hour),
	})

	if err != nil {
		request.ServerError(w, r, err, "req_id", reqID)
		return
	}

	logger.Debug("New session", "req_id", reqID)

	_, err = csrf.AddNewCSRFToken(session.ID)
	if err != nil {
		request.ServerError(w, r, err, "req_id", reqID)
		return
	}

	secure := config.Vars.GoatEnv != "dev"
	httponly := config.Vars.GoatEnv != "dev"

	cookie := &http.Cookie{
		Name:     "sessionToken",
		Value:    session.ID,
		Path:     "/",
		HttpOnly: httponly,
		SameSite: http.SameSiteLaxMode,
		Secure:   secure,
		MaxAge:   3600,
	}

	http.SetCookie(w, cookie)

	logger.Debug("Logged in", "req_id", reqID)
	request.HxRedirect(w, r, "/", "req_id", reqID)
}
