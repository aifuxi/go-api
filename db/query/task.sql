-- name: CreateTask :one
INSERT INTO "tasks"
(content, is_done)
VALUES ($1, $2)RETURNING *;

-- name: ListTasks :many
SELECT * FROM "tasks"
ORDER BY id LIMIT $1 OFFSET $2;