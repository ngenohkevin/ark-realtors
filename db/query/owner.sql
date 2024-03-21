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