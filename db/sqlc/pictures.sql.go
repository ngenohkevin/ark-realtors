// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: pictures.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const createPictures = `-- name: CreatePictures :one
INSERT INTO pictures (id, property_id, img_url, description)
VALUES ($1, $2, $3, $4) RETURNING id, property_id, img_url, description
`

type CreatePicturesParams struct {
	ID          uuid.UUID `json:"id"`
	PropertyID  uuid.UUID `json:"property_id"`
	ImgUrl      string    `json:"img_url"`
	Description string    `json:"description"`
}

func (q *Queries) CreatePictures(ctx context.Context, arg CreatePicturesParams) (Picture, error) {
	row := q.db.QueryRow(ctx, createPictures,
		arg.ID,
		arg.PropertyID,
		arg.ImgUrl,
		arg.Description,
	)
	var i Picture
	err := row.Scan(
		&i.ID,
		&i.PropertyID,
		&i.ImgUrl,
		&i.Description,
	)
	return i, err
}

const deletePictures = `-- name: DeletePictures :exec
DELETE FROM pictures WHERE id = $1
`

func (q *Queries) DeletePictures(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.Exec(ctx, deletePictures, id)
	return err
}

const getPictures = `-- name: GetPictures :one
SELECT id, property_id, img_url, description FROM pictures WHERE property_id = $1
`

func (q *Queries) GetPictures(ctx context.Context, propertyID uuid.UUID) (Picture, error) {
	row := q.db.QueryRow(ctx, getPictures, propertyID)
	var i Picture
	err := row.Scan(
		&i.ID,
		&i.PropertyID,
		&i.ImgUrl,
		&i.Description,
	)
	return i, err
}

const updatePictures = `-- name: UpdatePictures :exec
UPDATE pictures
SET img_url = $2, description = $3
WHERE id = $1
RETURNING id, property_id, img_url, description
`

type UpdatePicturesParams struct {
	ID          uuid.UUID `json:"id"`
	ImgUrl      string    `json:"img_url"`
	Description string    `json:"description"`
}

func (q *Queries) UpdatePictures(ctx context.Context, arg UpdatePicturesParams) error {
	_, err := q.db.Exec(ctx, updatePictures, arg.ID, arg.ImgUrl, arg.Description)
	return err
}