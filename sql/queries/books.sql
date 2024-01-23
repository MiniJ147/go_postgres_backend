-- name: CreateBook :one
INSERT INTO books(id, created_at, updated_at, title, author_id)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: FetchBooksByAuthorName :many
SELECT * FROM books WHERE author_id = ($1)
ORDER BY title;

-- name: FetchBooks :many
SELECT * FROM books;