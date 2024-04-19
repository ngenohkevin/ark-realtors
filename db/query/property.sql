-- name: CreateProperty :one
INSERT INTO property (id, type, price, status, bedroom, bathroom, location, size, contact)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING *;

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
    bedroom = $5,
    bathroom = $6,
    location = $7,
    size = $8,
    contact = $9
WHERE id = $1
RETURNING *;

-- name: DeleteProperty :exec
DELETE FROM property WHERE id = $1;