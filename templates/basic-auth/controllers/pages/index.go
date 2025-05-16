package pages

import (
	"basic-auth/controllers/helpers"
	"basic-auth/db/models"
	"basic-auth/views/components"
	"basic-auth/views/pages"
	"context"
	"errors"
	"log/slog"
	"net/http"

	"github.com/peterszarvas94/goat/pkg/csrf"
	"github.com/peterszarvas94/goat/pkg/database"
	"github.com/peterszarvas94/goat/pkg/request"
	"github.com/peterszarvas94/goat/pkg/server"
)

func Index(w http.ResponseWriter, r *http.Request) {
	reqID, ok := r.Context().Value("req_id").(string)
	if reqID == "" || !ok {
		request.ServerError(w, r, errors.New("Request ID is missing"))
		return
	}

	props := &pages.IndexProps{}

	ctxUser, ctxSession, err := helpers.CheckAuthStatus(r)
	if err != nil {
		slog.Debug(err.Error(), "req_id", reqID)
		slog.Debug("Rendering index page as guest", "req_id", reqID)
		server.Render(w, r, pages.Index(props), http.StatusOK)
		return
	}

	csrfToken, err := csrf.Get(ctxSession.ID)
	if err != nil {
		request.ServerError(w, r, err, "req_id", reqID)
		return
	}

	db, err := database.Get()
	if err != nil {
		request.ServerError(w, r, err, "req_id", reqID)
		return
	}

	queries := models.New(db)
	posts, err := queries.GetPostsByUserId(context.Background(), ctxUser.ID)
	if err != nil {
		request.ServerError(w, r, err, "req_id", reqID)
		return
	}

	props.UserinfoProps = &components.UserinfoProps{
		Name:  ctxUser.Name,
		Email: ctxUser.Email,
	}

	props.Posts = posts

	props.PostformProps = &components.PostformProps{
		CSRFToken: csrfToken,
		UserID:    ctxUser.ID,
	}

	slog.Debug("Rendering index page as user", "req_id", reqID)
	server.Render(w, r, pages.Index(props), http.StatusOK)
}
