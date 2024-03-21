-- name: CreateProperty :one
INSERT INTO property (id, type, price, status, pictures, bedroom, bathroom, location, size, contact, owner_id, agent_id)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12) RETURNING *;

-- name: GetProperty :one
SELECT * FROM property WHERE id = $1 LIMIT 1;

-- name: ListProperties :many
SELECT * FROM property
ORDER BY created_at DESC
LIMIT $1 OFFSET $2;

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
    contact = $10
WHERE id = $1
RETURNING *;

-- name: DeleteProperty :exec
DELETE FROM property WHERE id = $1;