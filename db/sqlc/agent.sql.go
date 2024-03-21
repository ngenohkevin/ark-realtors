// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: agent.sql

package db

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

const createAgent = `-- name: CreateAgent :one
INSERT INTO agent(id, phone_number, user_id, national_id, kra_pin)
VALUES ($1, $2, $3, $4, $5) RETURNING id, phone_number, user_id, national_id, kra_pin
`

func (q *Queries) CreateAgent(ctx context.Context, iD uuid.UUID, phoneNumber string, userID pgtype.UUID, nationalID string, kraPin string) (Agent, error) {
	row := q.db.QueryRow(ctx, createAgent,
		iD,
		phoneNumber,
		userID,
		nationalID,
		kraPin,
	)
	var i Agent
	err := row.Scan(
		&i.ID,
		&i.PhoneNumber,
		&i.UserID,
		&i.NationalID,
		&i.KraPin,
	)
	return i, err
}

const deleteAgent = `-- name: DeleteAgent :exec
DELETE FROM agent WHERE id = $1
`

func (q *Queries) DeleteAgent(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.Exec(ctx, deleteAgent, id)
	return err
}

const getAgent = `-- name: GetAgent :one
SELECT id, phone_number, user_id, national_id, kra_pin FROM agent WHERE id = $1 LIMIT 1
`

func (q *Queries) GetAgent(ctx context.Context, id uuid.UUID) (Agent, error) {
	row := q.db.QueryRow(ctx, getAgent, id)
	var i Agent
	err := row.Scan(
		&i.ID,
		&i.PhoneNumber,
		&i.UserID,
		&i.NationalID,
		&i.KraPin,
	)
	return i, err
}

const listAgents = `-- name: ListAgents :many
SELECT id, phone_number, user_id, national_id, kra_pin FROM agent
ORDER BY user_id DESC
`

func (q *Queries) ListAgents(ctx context.Context) ([]Agent, error) {
	rows, err := q.db.Query(ctx, listAgents)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Agent{}
	for rows.Next() {
		var i Agent
		if err := rows.Scan(
			&i.ID,
			&i.PhoneNumber,
			&i.UserID,
			&i.NationalID,
			&i.KraPin,
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

const updateAgent = `-- name: UpdateAgent :exec
UPDATE agent
SET phone_number = $2,
    national_id = $3,
    kra_pin = $4
WHERE id = $1
RETURNING id, phone_number, user_id, national_id, kra_pin
`

func (q *Queries) UpdateAgent(ctx context.Context, iD uuid.UUID, phoneNumber string, nationalID string, kraPin string) error {
	_, err := q.db.Exec(ctx, updateAgent,
		iD,
		phoneNumber,
		nationalID,
		kraPin,
	)
	return err
}
