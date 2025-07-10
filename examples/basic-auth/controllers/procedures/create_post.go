package procedures

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"basic-auth/controllers/helpers"
	"basic-auth/db/models"
	"basic-auth/views/components"

	"github.com/peterszarvas94/goat/pkg/database"
	"github.com/peterszarvas94/goat/pkg/server"
	"github.com/peterszarvas94/goat/pkg/uuid"
	"github.com/peterszarvas94/goat/pkg/validation"
)

type CreatePostRequest struct {
	Title   string `validate:"required,min=1,max=200"`
	Content string `validate:"required,min=1,max=5000"`
}

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

	title := r.FormValue("title")
	content := r.FormValue("content")

	postReq := CreatePostRequest{
		Title:   title,
		Content: content,
	}

	if err := validation.ValidateStruct(postReq); err != nil {
		messages := validation.BuildValidationMessages(err)
		helpers.BadRequest(w, r, messages, false, "req_id", reqID)
		return
	}

	db, err := database.Get()
	if err != nil {
		helpers.ServerError(w, r, []string{fmt.Sprintf("Can not get database: %s", err.Error())}, true, "req_id", reqID)
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
		helpers.ServerError(w, r, []string{fmt.Sprintf("Can not creat post: %s": err.Error())}, true, "req_id", reqID)
		return
	}

	slog.Debug("Post created, rendering new post", "req_id", reqID)

	server.Render(w, r, components.Post(&models.Post{
		ID:    post.ID,
		Title: post.Title,
	}), http.StatusOK)
}
