-- name: CreateUser :one
INSERT INTO users (id, username, full_name, email, hashed_password)
VALUES ($1, $2, $3, $4, $5) RETURNING *;

-- name: GetUser :one
SELECT * FROM users WHERE id = $1 LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1 LIMIT 1;

-- name: UpdateUser :exec
UPDATE users
SET username = $2,
    full_name = $3,
    email = $4,
    hashed_password = $5
WHERE id = $1
RETURNING *;

-- name CreateProperty :one
INSERT INTO property (id, type, price, status, pictures, bedroom, bathroom, location, size, contact, owner_id, agent_id)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12) RETURNING *;

-- name: GetProperty :one
SELECT * FROM property WHERE id = $1 LIMIT 1;

-- name: ListProperties :many
SELECT * FROM property
ORDER BY created_at DESC;

-- name: UpdateProperty :exec
UPDATE property
SET type = $2,
    price = $3,
    status = $4,
    pictures = $5,
    bedroom = $6,
    bathroom = $7,
    location = $8,
    size = $9,
    contact = $10,
    owner_id = $11,
    agent_id = $12
WHERE id = $1
RETURNING *;

-- name: DeleteProperty :exec
DELETE FROM property WHERE id = $1;

-- name: CreateAgent :one
INSERT INTO agent(id, phone_number, user_id, national_id, kra_pin)
VALUES ($1, $2, $3, $4, $5) RETURNING *;

-- name: GetAgent :one
SELECT * FROM agent WHERE id = $1 LIMIT 1;

-- name: ListAgents :many
SELECT * FROM agent
ORDER BY user_id DESC;

-- name: UpdateAgent :exec
UPDATE agent
SET phone_number = $2,
    national_id = $3,
    kra_pin = $4
WHERE id = $1
RETURNING *;

-- name: DeleteAgent :exec
DELETE FROM agent WHERE id = $1;

-- name: CreateOwner :one
INSERT INTO owner(id, phone_number, user_id, national_id)
VALUES ($1, $2, $3, $4) RETURNING *;

-- name: GetOwner :one
SELECT * FROM owner WHERE id = $1 LIMIT 1;

-- name: ListOwners :many
SELECT * FROM owner
ORDER BY user_id DESC;

-- name: UpdateOwner :exec
UPDATE owner
SET phone_number = $2,
    national_id = $3
WHERE id = $1
RETURNING *;

-- name: DeleteOwner :exec
DELETE FROM owner WHERE id = $1;

