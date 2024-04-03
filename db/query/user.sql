-- name: CreateUser :one
INSERT INTO users (id, username, full_name, email, hashed_password, role)
VALUES ($1, $2, $3, $4, $5, $6) RETURNING *;

-- name: GetUser :one
SELECT * FROM users WHERE username = $1 LIMIT 1;

-- name: UpdateUser :exec
UPDATE users
SET username = $2,
    full_name = $3,
    email = $4,
    hashed_password = $5,
    role = $6
WHERE id = $1
RETURNING *;

