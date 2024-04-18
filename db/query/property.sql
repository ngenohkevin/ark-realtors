-- name: CreateProperty :one
INSERT INTO property (id, type, price, status, img_url, bedroom, bathroom, location, size, contact)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING *;

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
    img_url = $5,
    bedroom = $6,
    bathroom = $7,
    location = $8,
    size = $9,
    contact = $10
WHERE id = $1
RETURNING *;

-- name: DeleteProperty :exec
DELETE FROM property WHERE id = $1;