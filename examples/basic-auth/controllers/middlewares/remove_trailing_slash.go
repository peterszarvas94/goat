package middlewares

import (
	"net/http"
	"strings"
)

func RemoveTrailingSlash(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" && strings.HasSuffix(r.URL.Path, "/") {
			http.Redirect(w, r, r.URL.Path[:len(r.URL.Path)-1], http.StatusMovedPermanently)
			return
		}
		next.ServeHTTP(w, r)
	}
}
