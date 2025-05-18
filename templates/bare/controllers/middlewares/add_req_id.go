package middlewares

import (
	"context"
	"net/http"

	"github.com/peterszarvas94/goat/pkg/uuid"
)

func AddRequestId(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reqID := uuid.New("req")

		ctx := context.WithValue(r.Context(), "req_id", reqID)
		next(w, r.WithContext(ctx))
	}
}
