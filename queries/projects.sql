-- name: ListProjects :many
SELECT *
FROM projects
WHERE user_id = ?
ORDER BY name ASC, created_at DESC;

-- name: GetProject :one
SELECT *
FROM projects
WHERE id = ?
AND user_id = ?;

-- name: CreateProject :one
INSERT INTO projects (
  user_id,
  name
) VALUES (?, ?)
RETURNING *;

-- name: DeleteProject :exec
DELETE FROM projects
WHERE id = ?
AND user_id = ?;
