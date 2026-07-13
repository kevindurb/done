-- name: CreateTask :one
INSERT INTO tasks (
  user_id,
  project_id,
  description
) VALUES (?, ?, ?)
RETURNING *;

-- name: ListTasks :many
SELECT *
FROM tasks
WHERE user_id = ?
AND done = FALSE;

-- name: ListTasksByProject :many
SELECT *
FROM tasks
WHERE user_id = ?
AND project_id = ?
AND done = FALSE;

-- name: ListTasksDone :many
SELECT *
FROM tasks
WHERE user_id = ?
AND done = TRUE;

-- name: ListTasksDoneByProject :many
SELECT *
FROM tasks
WHERE user_id = ?
AND project_id = ?
AND done = TRUE;

-- name: GetTask :one
SELECT *
FROM tasks
WHERE user_id = ?
AND id = ?;

-- name: MarkTaskDone :one
UPDATE tasks
SET done = TRUE
WHERE id = ?
AND user_id = ?
RETURNING *;
