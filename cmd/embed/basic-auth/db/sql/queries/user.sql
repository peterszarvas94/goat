-- name: GetUserByID :one
SELECT *
FROM user
WHERE id = ?;

-- name: GetUserByEmail :one
SELECT *
FROM user
WHERE email = ?;

-- name: ListUser :many
SELECT *
FROM user
ORDER BY name;

-- name: CreateUser :one
INSERT INTO user (id, name, email, password)
VALUES (?, ?, ?, ?)
RETURNING id, name, email;

-- name: UpdateUser :one
UPDATE user
SET name = ?, email = ?, password = ?
WHERE id = ?
RETURNING id, name, email;

-- name: DeleteUser :exec
DELETE FROM user
WHERE id = ?;
