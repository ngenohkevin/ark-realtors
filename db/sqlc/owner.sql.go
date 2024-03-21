// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: owner.sql

package db

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

const createOwner = `-- name: CreateOwner :one
INSERT INTO owner(id, phone_number, user_id, national_id)
VALUES ($1, $2, $3, $4) RETURNING id, phone_number, user_id, national_id
`

func (q *Queries) CreateOwner(ctx context.Context, iD uuid.UUID, phoneNumber string, userID pgtype.UUID, nationalID string) (Owner, error) {
	row := q.db.QueryRow(ctx, createOwner,
		iD,
		phoneNumber,
		userID,
		nationalID,
	)
	var i Owner
	err := row.Scan(
		&i.ID,
		&i.PhoneNumber,
		&i.UserID,
		&i.NationalID,
	)
	return i, err
}

const deleteOwner = `-- name: DeleteOwner :exec
DELETE FROM owner WHERE id = $1
`

func (q *Queries) DeleteOwner(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.Exec(ctx, deleteOwner, id)
	return err
}

const getOwner = `-- name: GetOwner :one
SELECT id, phone_number, user_id, national_id FROM owner WHERE id = $1 LIMIT 1
`

func (q *Queries) GetOwner(ctx context.Context, id uuid.UUID) (Owner, error) {
	row := q.db.QueryRow(ctx, getOwner, id)
	var i Owner
	err := row.Scan(
		&i.ID,
		&i.PhoneNumber,
		&i.UserID,
		&i.NationalID,
	)
	return i, err
}

const listOwners = `-- name: ListOwners :many
SELECT id, phone_number, user_id, national_id FROM owner
ORDER BY user_id DESC
`

func (q *Queries) ListOwners(ctx context.Context) ([]Owner, error) {
	rows, err := q.db.Query(ctx, listOwners)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Owner{}
	for rows.Next() {
		var i Owner
		if err := rows.Scan(
			&i.ID,
			&i.PhoneNumber,
			&i.UserID,
			&i.NationalID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateOwner = `-- name: UpdateOwner :exec
UPDATE owner
SET phone_number = $2,
    national_id = $3
WHERE id = $1
RETURNING id, phone_number, user_id, national_id
`

func (q *Queries) UpdateOwner(ctx context.Context, iD uuid.UUID, phoneNumber string, nationalID string) error {
	_, err := q.db.Exec(ctx, updateOwner, iD, phoneNumber, nationalID)
	return err
}