package pages

import (
	"basic-auth/db/models"
	"basic-auth/views/pages"
	"context"
	"errors"
	"log/slog"
	"net/http"

	"github.com/peterszarvas94/goat/pkg/database"
	"github.com/peterszarvas94/goat/pkg/request"
	"github.com/peterszarvas94/goat/pkg/server"
)

func PostPageHandler(w http.ResponseWriter, r *http.Request) {
	reqID, ok := r.Context().Value("req_id").(string)
	if reqID == "" || !ok {
		request.ServerError(w, r, errors.New("Request ID is missing"))
		return
	}

	postId := r.PathValue("id")

	db, err := database.Get()
	if err != nil {
		request.ServerError(w, r, err, "req_id", reqID)
		return
	}

	queries := models.New(db)
	post, err := queries.GetPostByID(context.Background(), postId)
	if err != nil {
		request.ServerError(w, r, err, "req_id", reqID)
		return
	}

	slog.Debug("Rendering post page", "req_id", reqID)
	server.Render(w, r, pages.PostPageTemplate(&post), http.StatusOK)
}
