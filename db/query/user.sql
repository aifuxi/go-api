-- name: CreateUser :one
INSERT INTO "users"
  (username, hashed_password)
VALUES ($1, $2)RETURNING *;

-- name: ListUsers :many
SELECT * FROM "users"
ORDER BY id LIMIT $1 OFFSET $2;

-- name: GetUser :one
SELECT * FROM "users" WHERE username = $1;