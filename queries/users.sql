-- name: CreateUser :one
INSERT INTO users (
  email,
  hash
) VALUES (?, ?)
RETURNING *;

-- name: GetUserByEmail :one
SELECT *
FROM users
WHERE email = ?;
