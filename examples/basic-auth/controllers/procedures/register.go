package procedures

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"basic-auth/controllers/helpers"
	"basic-auth/db/models"

	"github.com/peterszarvas94/goat/pkg/database"
	"github.com/peterszarvas94/goat/pkg/hash"
	"github.com/peterszarvas94/goat/pkg/request"
	"github.com/peterszarvas94/goat/pkg/uuid"
	"github.com/peterszarvas94/goat/pkg/validation"
)

type UserRegistration struct {
	Username string `validate:"required,min=3,max=20,alphanum"`
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=8"`
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	reqID, ok := r.Context().Value("req_id").(string)
	if reqID == "" || !ok {
		helpers.ServerError(w, r, []string{"Request ID is missing"}, true)
		return
	}

	name := r.FormValue("name")
	email := r.FormValue("email")
	password := r.FormValue("password")

	userReg := UserRegistration{
		Username: name,
		Email:    email,
		Password: password,
	}

	if err := validation.ValidateStruct(userReg); err != nil {
		messages := validation.BuildValidationMessages(err)
		helpers.BadRequest(w, r, messages, false, "req_id", reqID)
		return
	}

	hashed, err := hash.HashPassword(password)
	if err != nil {
		helpers.ServerError(w, r, []string{fmt.Sprintf("Can not hash password: %s", err.Error())}, true, "req_id", reqID)
		return
	}

	db, err := database.Get()
	if err != nil {
		helpers.ServerError(w, r, []string{fmt.Sprintf("Can not get database: %s", err.Error())}, true, "req_id", reqID)
		return
	}

	queries := models.New(db)

	conflictMessages := []string{}

	existing, err := queries.GetUserByEmail(context.Background(), email)
	if err == nil || existing.Email == email {
		conflictMessages = append(conflictMessages, "Email already in use")
	}

	existing, err = queries.GetUserByName(context.Background(), name)
	if err == nil || existing.Email == email {
		conflictMessages = append(conflictMessages, "Name already in use")
	}

	if len(conflictMessages) > 0 {
		helpers.Conflict(w, r, conflictMessages, false, "req_id", reqID)
		return
	}

	_, err = queries.CreateUser(context.Background(), models.CreateUserParams{
		ID:       uuid.New("usr"),
		Name:     name,
		Email:    email,
		Password: hashed,
	})

	if err != nil {
		helpers.ServerError(w, r, []string{fmt.Sprintf("Can not create user: %s", err.Error())}, true, "req_id", reqID)
		return
	}

	slog.Debug("Registered", "req_id", reqID)
	request.HxRedirect(w, r, "/login", "req_id", reqID)
}
