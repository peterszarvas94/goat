package helpers

import (
	"errors"
	"net/http"
	"scaffhold/db/models"
)

func CheckAuthStatus(r *http.Request) (*models.User, *models.Session, error) {
	session, ok := r.Context().Value("session").(*models.Session)
	if !ok || session == nil {
		return nil, nil, errors.New("Session is missing")
	}

	user, ok := r.Context().Value("user").(*models.User)
	if !ok || user == nil {
		return nil, nil, errors.New("User is missing")
	}

	return user, session, nil
}
