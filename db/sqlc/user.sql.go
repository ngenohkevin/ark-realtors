// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: user.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (id, username, full_name, email, hashed_password)
VALUES ($1, $2, $3, $4, $5) RETURNING id, username, full_name, email, hashed_password, password_changed_at, created_at
`

func (q *Queries) CreateUser(ctx context.Context, iD uuid.UUID, username string, fullName string, email string, hashedPassword string) (User, error) {
	row := q.db.QueryRow(ctx, createUser,
		iD,
		username,
		fullName,
		email,
		hashedPassword,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.FullName,
		&i.Email,
		&i.HashedPassword,
		&i.PasswordChangedAt,
		&i.CreatedAt,
	)
	return i, err
}

const getUser = `-- name: GetUser :one
SELECT id, username, full_name, email, hashed_password, password_changed_at, created_at FROM users WHERE id = $1 LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, id uuid.UUID) (User, error) {
	row := q.db.QueryRow(ctx, getUser, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.FullName,
		&i.Email,
		&i.HashedPassword,
		&i.PasswordChangedAt,
		&i.CreatedAt,
	)
	return i, err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, username, full_name, email, hashed_password, password_changed_at, created_at FROM users WHERE email = $1 LIMIT 1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRow(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.FullName,
		&i.Email,
		&i.HashedPassword,
		&i.PasswordChangedAt,
		&i.CreatedAt,
	)
	return i, err
}

const updateUser = `-- name: UpdateUser :exec
UPDATE users
SET username = $2,
    full_name = $3,
    email = $4,
    hashed_password = $5
WHERE id = $1
RETURNING id, username, full_name, email, hashed_password, password_changed_at, created_at
`

func (q *Queries) UpdateUser(ctx context.Context, iD uuid.UUID, username string, fullName string, email string, hashedPassword string) error {
	_, err := q.db.Exec(ctx, updateUser,
		iD,
		username,
		fullName,
		email,
		hashedPassword,
	)
	return err
}