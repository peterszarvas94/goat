package procedures

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"basic-auth/config"
	"basic-auth/controllers/helpers"
	"basic-auth/db/models"

	"github.com/peterszarvas94/goat/pkg/csrf"
	"github.com/peterszarvas94/goat/pkg/database"
	"github.com/peterszarvas94/goat/pkg/hash"
	"github.com/peterszarvas94/goat/pkg/request"
	"github.com/peterszarvas94/goat/pkg/uuid"
	"github.com/peterszarvas94/goat/pkg/validation"
)

type LoginRequest struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	reqID, ok := r.Context().Value("req_id").(string)
	if reqID == "" || !ok {
		helpers.ServerError(w, r, []string{"Request ID is missing"}, true)
		return
	}

	email := r.FormValue("email")
	password := r.FormValue("password")

	loginReq := LoginRequest{
		Email:    email,
		Password: password,
	}

	if err := validation.ValidateStruct(loginReq); err != nil {
		messages := validation.BuildValidationMessages(err)
		helpers.BadRequest(w, r, messages, false, "req_id", reqID)
		return
	}

	db, err := database.Get()
	if err != nil {
		helpers.ServerError(w, r, []string{fmt.Sprintf("Can not get database: %s", err.Error())}, true, "req_id", reqID)
		return
	}

	queries := models.New(db)
	user, err := queries.GetUserByEmail(context.Background(), email)
	if err != nil {
		helpers.Unauthorized(w, r, []string{"User with this email not found"}, true, "req_id", reqID)
		return
	}

	valid := hash.VerifyPassword(password, user.Password)
	if !valid {
		helpers.Unauthorized(w, r, []string{"Bad credentials"}, true, "req_id", reqID)
		return
	}

	slog.Debug("Credentials are valid", "req_id", reqID)

	sessionId := uuid.New("ses")
	session, err := queries.CreateSession(context.Background(), models.CreateSessionParams{
		ID:         sessionId,
		UserID:     user.ID,
		ValidUntil: time.Now().Add(24 * time.Hour),
	})

	if err != nil {
		helpers.ServerError(w, r, []string{fmt.Sprintf("Can not create session", err.Error())}, true, "req_id", reqID)
		return
	}

	slog.Debug("New session", "req_id", reqID)

	_, err = csrf.AddNewCSRFToken(session.ID)
	if err != nil {
		helpers.ServerError(w, r, []string{fmt.Sprintf("Can not add new csrf token: %s", err.Error())}, true, "req_id", reqID)
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

	slog.Debug("Logged in", "req_id", reqID)
	request.HxRedirect(w, r, "/", "req_id", reqID)
}
