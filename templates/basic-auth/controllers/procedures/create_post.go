package procedures

import (
	"basic-auth/db/models"
	"basic-auth/views/components"
	"context"
	"errors"
	"log/slog"
	"net/http"

	"github.com/peterszarvas94/goat/pkg/database"
	"github.com/peterszarvas94/goat/pkg/request"
	"github.com/peterszarvas94/goat/pkg/server"
	"github.com/peterszarvas94/goat/pkg/uuid"
)

func CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	reqID, ok := r.Context().Value("req_id").(string)
	if reqID == "" || !ok {
		request.ServerError(w, r, errors.New("Request ID is missing"))
		return
	}

	ctxUser, ok := r.Context().Value("user").(*models.User)
	if ctxUser == nil || !ok {
		request.ServerError(w, r, errors.New("User is missing"), "req_id", reqID)
		return
	}

	if err := r.ParseForm(); err != nil {
		request.ServerError(w, r, err, "req_id", reqID)
		return
	}

	title := r.FormValue("title")
	if title == "" {
		request.HxReswap(w, "innerHTML")
		request.BadRequest(w, r, errors.New("Title can not be empty"), "req_id", reqID)
		return
	}

	content := r.FormValue("content")
	if content == "" {
		request.HxReswap(w, "innerHTML")
		request.BadRequest(w, r, errors.New("Content can not be empty"), "req_id", reqID)
		return
	}

	db, err := database.Get()
	if err != nil {
		request.ServerError(w, r, err, "req_id", reqID)
		return
	}

	postId := uuid.New("pst")
	queries := models.New(db)
	post, err := queries.CreatePost(context.Background(), models.CreatePostParams{
		ID:      postId,
		Title:   title,
		Content: content,
		UserID:  ctxUser.ID,
	})

	if err != nil {
		request.ServerError(w, r, err, "req_id", reqID)
		return
	}

	slog.Debug("Post created, rendering new post", "req_id", reqID)

	server.Render(w, r, components.Post(&models.Post{
		ID:    post.ID,
		Title: post.Title,
	}), http.StatusOK)
}
