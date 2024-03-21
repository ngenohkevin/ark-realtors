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