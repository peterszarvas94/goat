package middlewares

import (
	"net/http"
	"scaffhold/config"
)

func Cache(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if config.Vars.GoatEnv == "dev" && r.URL.Path != "/favicon.ico" {
			w.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate, max-age=0")
			w.Header().Set("Pragma", "no-cache")
			w.Header().Set("Expires", "0")
		}

		next(w, r)
	}
}
