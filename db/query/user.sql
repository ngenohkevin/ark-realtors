-- name: CreateUser :one
INSERT INTO users (id, username, full_name, email, hashed_password)
VALUES ($1, $2, $3, $4, $5) RETURNING *;

-- name: GetUser :one
SELECT * FROM users WHERE username = $1 LIMIT 1;

-- name: UpdateUser :one
UPDATE users
SET username = COALESCE(sqlc.narg(username), username),
    full_name = COALESCE(sqlc.narg(full_name), full_name),
    email = COALESCE(sqlc.narg(email), email),
    hashed_password = COALESCE(sqlc.narg(hashed_password), hashed_password),
    role = COALESCE(sqlc.narg(role), role)
WHERE id = sqlc.arg(id)
RETURNING *;

