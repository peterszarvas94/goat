package procedures

import (
	"basic-auth/controllers/helpers"
	"basic-auth/db/models"
	"basic-auth/views/components"
	"context"
	"log/slog"
	"net/http"

	"github.com/peterszarvas94/goat/pkg/database"
	"github.com/peterszarvas94/goat/pkg/server"
	"github.com/peterszarvas94/goat/pkg/uuid"
)

func CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	reqID, ok := r.Context().Value("req_id").(string)
	if reqID == "" || !ok {
		helpers.ServerError(w, r, []string{"Request ID is missing"}, true)
		return
	}

	ctxUser, ok := r.Context().Value("user").(*models.User)
	if ctxUser == nil || !ok {
		helpers.ServerError(w, r, []string{"User is missing"}, true, "req_id", reqID)
		return
	}

	if err := r.ParseForm(); err != nil {
		helpers.ServerError(w, r, []string{err.Error()}, true, "req_id", reqID)
		return
	}

	messages := []string{}

	title := r.FormValue("title")
	if title == "" {
		messages = append(messages, "Title can not be empty")
	}

	content := r.FormValue("content")
	if content == "" {
		messages = append(messages, "Content can not be empty")
	}

	if len(messages) > 0 {
		helpers.BadRequest(w, r, messages, false, "req_id", reqID)
		return
	}

	db, err := database.Get()
	if err != nil {
		helpers.ServerError(w, r, []string{err.Error()}, true, "req_id", reqID)
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
		helpers.ServerError(w, r, []string{err.Error()}, true, "req_id", reqID)
		return
	}

	slog.Debug("Post created, rendering new post", "req_id", reqID)

	server.Render(w, r, components.Post(&models.Post{
		ID:    post.ID,
		Title: post.Title,
	}), http.StatusOK)
}
