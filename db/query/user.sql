-- name: CreateUser :one
INSERT INTO users (first_name,
                   last_name,
                   age)
VALUES ($1, $2, $3)
RETURNING *;

-- name: DeleteUser :exec
DELETE
FROM users
WHERE id = $1;

-- name: GetUser :one
SELECT *
FROM users
WHERE id = $1
LIMIT 1;

-- name: GetUserForUpdate :one
SELECT *
FROM users
WHERE id = $1
LIMIT 1 FOR NO KEY UPDATE;

-- name: ListUsers :many
SELECT *
FROM users
ORDER BY id
LIMIT $1 OFFSET $2;

-- name: UpdateUser :one
UPDATE users
SET first_name = $2,
    last_name = $3,
    age = $4
WHERE id = $1
RETURNING *;