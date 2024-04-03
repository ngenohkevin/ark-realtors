// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	CreateProperty(ctx context.Context, arg CreatePropertyParams) (Property, error)
	CreateSession(ctx context.Context, arg CreateSessionParams) (Session, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	DeleteProperty(ctx context.Context, id uuid.UUID) error
	GetProperty(ctx context.Context, id uuid.UUID) (Property, error)
	GetSession(ctx context.Context, id uuid.UUID) (Session, error)
	GetUser(ctx context.Context, username string) (User, error)
	ListProperties(ctx context.Context, arg ListPropertiesParams) ([]Property, error)
	UpdateProperty(ctx context.Context, arg UpdatePropertyParams) error
	UpdateUser(ctx context.Context, arg UpdateUserParams) error
}

var _ Querier = (*Queries)(nil)
