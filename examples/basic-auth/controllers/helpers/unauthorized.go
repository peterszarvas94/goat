package helpers

import (
	"basic-auth/views/components"
	"log/slog"
	"net/http"
	"slices"

	"github.com/peterszarvas94/goat/pkg/request"
	"github.com/peterszarvas94/goat/pkg/server"
)

func Unauthorized(w http.ResponseWriter, r *http.Request, messages []string, hide bool, args ...any) {
	toastMessages := []components.ToastMessage{}

	for message := range slices.Values(messages) {
		slog.Error(message, args...)

		m := message
		if hide {
			m = "Unauthorized"
		}

		toastMessages = append(toastMessages, components.ToastMessage{
			Message: m,
			Level:   "error",
		})
	}

	request.HxReswap(w, "innerHTML")
	server.Render(w, r, components.Toast(toastMessages), http.StatusUnauthorized)
}
