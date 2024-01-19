-- name: CreateBook :one
INSERT INTO books(id, created_at, updated_at, title, author_id)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;