-- name: GetAuthor :one
SELECT * FROM authors WHERE id = ?;


-- name: ListAuthors :many
SELECT * FROM authors ORDER BY id;


-- name: CreateAuthor :execresult
INSERT INTO authors (name, country, age) VALUES (?, ?, ?);


-- name: UpdateAuthor :execresult
UPDATE authors SET name = ?, country = ?, age = ? WHERE id = ?;


-- name: DeleteAuthor :exec
DELETE FROM authors WHERE id = ?;