-- name: CreateUser :one
INSERT INTO users (
  email,
  hash
) VALUES (?, ?)
RETURNING *;
