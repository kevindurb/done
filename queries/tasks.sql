-- name: CreateTask :one
INSERT INTO tasks (
  user_id,
  description
) VALUES (?, ?)
RETURNING *;

-- name: ListTasks :many
SELECT *
FROM tasks
WHERE user_id = ?
AND done = FALSE;
