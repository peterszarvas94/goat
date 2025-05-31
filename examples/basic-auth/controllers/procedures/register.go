package procedures

import (
	"context"
	"errors"
	"log/slog"
	"net/http"

	"basic-auth/db/models"

	"github.com/peterszarvas94/goat/pkg/database"
	"github.com/peterszarvas94/goat/pkg/hash"
	"github.com/peterszarvas94/goat/pkg/request"
	"github.com/peterszarvas94/goat/pkg/uuid"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	reqID, ok := r.Context().Value("req_id").(string)
	if reqID == "" || !ok {
		request.ServerError(w, r, errors.New("Request ID is missing"))
		return
	}

	name := r.FormValue("name")
	if name == "" {
		request.BadRequest(w, r, errors.New("Name can not be empty"), "req_id", reqID)
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

	hashed, err := hash.HashPassword(password)
	if err != nil {
		request.ServerError(w, r, err, "req_id", reqID)
		return
	}

	db, err := database.Get()
	if err != nil {
		request.ServerError(w, r, err, "req_id", reqID)
		return
	}

	queries := models.New(db)

	existing, err := queries.GetUserByEmail(context.Background(), email)
	// user conflict
	if err == nil {
		if existing.Name == name {
			request.Conflict(w, r, errors.New("Name already in use"), "req_id", reqID)
			return
		}

		if existing.Email == email {
			request.Conflict(w, r, errors.New("Email already in use"), "req_id", reqID)
			return
		}
	}

	_, err = queries.CreateUser(context.Background(), models.CreateUserParams{
		ID:       uuid.New("usr"),
		Name:     name,
		Email:    email,
		Password: hashed,
	})

	if err != nil {
		request.ServerError(w, r, err, "req_id", reqID)
		return
	}

	slog.Debug("Registered", "req_id", reqID)
	request.HxRedirect(w, r, "/login", "req_id", reqID)
}
