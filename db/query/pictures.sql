-- name: CreatePictures :one
INSERT INTO pictures (id, property_id, img_url, description)
VALUES ($1, $2, $3, $4) RETURNING *;

-- name: GetPictures :one
SELECT * FROM pictures WHERE property_id = $1;

-- name: UpdatePictures :exec
UPDATE pictures
SET img_url = $2, description = $3
WHERE id = $1
RETURNING *;

-- name: DeletePictures :exec
DELETE FROM pictures WHERE id = $1;