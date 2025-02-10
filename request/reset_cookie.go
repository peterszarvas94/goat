package request

import (
	"net/http"
)

func ResetCookie(w *http.ResponseWriter, cookieName string) {
	cookie := &http.Cookie{
		Name:   cookieName,
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}

	http.SetCookie(*w, cookie)
}
