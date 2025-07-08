package procedures

import (
	"context"
	"log/slog"
	"net/http"

	"basic-auth/controllers/helpers"
	"basic-auth/db/models"

	"github.com/peterszarvas94/goat/pkg/database"
	"github.com/peterszarvas94/goat/pkg/hash"
	"github.com/peterszarvas94/goat/pkg/request"
	"github.com/peterszarvas94/goat/pkg/uuid"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	reqID, ok := r.Context().Value("req_id").(string)
	if reqID == "" || !ok {
		helpers.ServerError(w, r, []string{"Request ID is missing"}, true)
		return
	}

	messages := []string{}

	name := r.FormValue("name")
	if name == "" {
		messages = append(messages, "Name can not be empty")
	}

	email := r.FormValue("email")
	if email == "" {
		messages = append(messages, "Email can not be empty")
	}

	password := r.FormValue("password")
	if password == "" {
		messages = append(messages, "Password can not be empty")
	}

	if len(messages) > 0 {
		helpers.BadRequest(w, r, messages, false, "req_id", reqID)
		return
	}

	hashed, err := hash.HashPassword(password)
	if err != nil {
		helpers.ServerError(w, r, []string{err.Error()}, true, "req_id", reqID)
		return
	}

	db, err := database.Get()
	if err != nil {
		helpers.ServerError(w, r, []string{err.Error()}, true, "req_id", reqID)
		return
	}

	queries := models.New(db)

	messages = []string{}

	existing, err := queries.GetUserByEmail(context.Background(), email)
	// user conflict
	if err == nil {
		if existing.Name == name {
			messages = append(messages, "Name already in use")
		}

		if existing.Email == email {
			messages = append(messages, "Email already in use")
		}
	}

	if len(messages) > 0 {
		helpers.BadRequest(w, r, messages, false, "req_id", reqID)
		return
	}

	_, err = queries.CreateUser(context.Background(), models.CreateUserParams{
		ID:       uuid.New("usr"),
		Name:     name,
		Email:    email,
		Password: hashed,
	})

	if err != nil {
		helpers.ServerError(w, r, []string{err.Error()}, true, "req_id", reqID)
		return
	}

	slog.Debug("Registered", "req_id", reqID)
	request.HxRedirect(w, r, "/login", "req_id", reqID)
}
