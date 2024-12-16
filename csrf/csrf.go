package csrf

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"sync"
)

var csrfTokens = sync.Map{}

func generateCSRFToken() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	hash := sha256.Sum256(b)
	return base64.URLEncoding.EncodeToString(hash[:]), nil
}

func AddNewCSRFToken(sessionID string) (string, error) {
	if _, ok := csrfTokens.Load(sessionID); ok {
		return "", fmt.Errorf("CSRF token does not exist for session \"%s\"", sessionID)
	}

	newToken, err := generateCSRFToken()
	if err != nil {
		return "", err
	}

	csrfTokens.Store(sessionID, newToken)

	return newToken, nil
}

func GetCSRFToken(sessionID string) (string, error) {
	value, ok := csrfTokens.Load(sessionID)
	if !ok {
		return "", fmt.Errorf("CSRF token does not exist for session \"%s\"", sessionID)
	}

	storedToken, ok := value.(string)
	if !ok {
		return "", fmt.Errorf("Stored CSRF token is not a valid string for session \"%s\"", sessionID)
	}

	return storedToken, nil
}

func ValidateCSRFToken(sessionID, csrfToken string) error {
	value, ok := csrfTokens.Load(sessionID)

	if !ok {
		return fmt.Errorf("CSRF token does not exist for session \"%s\"", sessionID)
	}

	storedToken, ok := value.(string)
	if !ok {
		return errors.New("Stored CSRF token is not a valid string")
	}

	if storedToken != csrfToken {
		return errors.New("CSRF token is not valid")
	}

	return nil
}

func DeleteCSRFToken(sessionId string) {
	csrfTokens.Delete(sessionId)
}
