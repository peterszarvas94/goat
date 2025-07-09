package middleware

import (
	"net/http"

	"github.com/unrolled/secure"
)

func SecurityHeaders(isDevelopment bool) func(http.HandlerFunc) http.HandlerFunc {
	secureMiddleware := secure.New(secure.Options{
		AllowedHosts:            []string{}, // Add your domains in production
		SSLRedirect:             !isDevelopment,
		SSLHost:                 "", // Your SSL host
		STSSeconds:              0,  // Disable HSTS in development
		STSIncludeSubdomains:    !isDevelopment,
		STSPreload:              !isDevelopment,
		FrameDeny:               true,
		ContentTypeNosniff:      true,
		BrowserXssFilter:        true,
		ContentSecurityPolicy:   "default-src 'self'; script-src 'self' 'unsafe-inline'; style-src 'self' 'unsafe-inline'; img-src 'self' data:; font-src 'self'; connect-src 'self'",
		ReferrerPolicy:          "strict-origin-when-cross-origin",
		CustomFrameOptionsValue: "DENY",
	})

	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			err := secureMiddleware.Process(w, r)
			if err != nil {
				return
			}
			next(w, r)
		}
	}
}
