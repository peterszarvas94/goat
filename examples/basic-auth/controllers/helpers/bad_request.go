package helpers

import (
	"basic-auth/views/components"
	"net/http"
	"slices"

	"log/slog"

	"github.com/peterszarvas94/goat/pkg/request"
	"github.com/peterszarvas94/goat/pkg/server"
)

func BadRequest(w http.ResponseWriter, r *http.Request, messages []string, hide bool, args ...any) {

	toastMessages := []components.ToastMessage{}

	for message := range slices.Values(messages) {
		slog.Error(message, args...)

		m := message
		if hide {
			m = "Bad request"
		}

		toastMessages = append(toastMessages, components.ToastMessage{
			Message: m,
			Level:   "error",
		})
	}

	request.HxReswap(w, "innerHTML")
	server.Render(w, r, components.Toast(toastMessages), http.StatusBadRequest)
}
