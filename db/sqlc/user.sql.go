// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: user.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (id, username, full_name, email, hashed_password, role)
VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, username, full_name, email, hashed_password, role, password_changed_at, created_at
`

type CreateUserParams struct {
	ID             uuid.UUID `json:"id"`
	Username       string    `json:"username"`
	FullName       string    `json:"full_name"`
	Email          string    `json:"email"`
	HashedPassword string    `json:"hashed_password"`
	Role           string    `json:"role"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, createUser,
		arg.ID,
		arg.Username,
		arg.FullName,
		arg.Email,
		arg.HashedPassword,
		arg.Role,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.FullName,
		&i.Email,
		&i.HashedPassword,
		&i.Role,
		&i.PasswordChangedAt,
		&i.CreatedAt,
	)
	return i, err
}

const getUser = `-- name: GetUser :one
SELECT id, username, full_name, email, hashed_password, role, password_changed_at, created_at FROM users WHERE username = $1 LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRow(ctx, getUser, username)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.FullName,
		&i.Email,
		&i.HashedPassword,
		&i.Role,
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
    hashed_password = $5,
    role = $6
WHERE id = $1
RETURNING id, username, full_name, email, hashed_password, role, password_changed_at, created_at
`

type UpdateUserParams struct {
	ID             uuid.UUID `json:"id"`
	Username       string    `json:"username"`
	FullName       string    `json:"full_name"`
	Email          string    `json:"email"`
	HashedPassword string    `json:"hashed_password"`
	Role           string    `json:"role"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) error {
	_, err := q.db.Exec(ctx, updateUser,
		arg.ID,
		arg.Username,
		arg.FullName,
		arg.Email,
		arg.HashedPassword,
		arg.Role,
	)
	return err
}
