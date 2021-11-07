-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY email;

-- name: CreateUser :one
INSERT INTO users (
    email, password
) VALUES (
             $1, $2
         )
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;