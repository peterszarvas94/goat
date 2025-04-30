-- name: GetPostByID :one
SELECT *
FROM post
WHERE id = ?;

-- name: ListPost :many
SELECT *
FROM post
ORDER BY name;

-- name: GetPostsByUserId :many
SELECT *
FROM post
WHERE user_id = ?
ORDER BY created_at;

-- name: CreatePost :one
INSERT INTO post (id, title, content, user_id)
VALUES (?, ?, ?, ?)
RETURNING id, title, content, user_id;

-- name: UpdatePost :one
UPDATE post
SET title = ?, content = ?, user_id = ?
WHERE id = ?
RETURNING id, title, content, user_id;

-- name: DeletePost :exec
DELETE FROM post
WHERE id = ?;
