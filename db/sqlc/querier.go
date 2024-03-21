// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type Querier interface {
	CreateAgent(ctx context.Context, iD uuid.UUID, phoneNumber string, userID pgtype.UUID, nationalID string, kraPin string) (Agent, error)
	CreateOwner(ctx context.Context, iD uuid.UUID, phoneNumber string, userID pgtype.UUID, nationalID string) (Owner, error)
	CreateProperty(ctx context.Context, arg CreatePropertyParams) (Property, error)
	CreateUser(ctx context.Context, iD uuid.UUID, username string, fullName string, email string, hashedPassword string) (User, error)
	DeleteAgent(ctx context.Context, id uuid.UUID) error
	DeleteOwner(ctx context.Context, id uuid.UUID) error
	DeleteProperty(ctx context.Context, id uuid.UUID) error
	GetAgent(ctx context.Context, id uuid.UUID) (Agent, error)
	GetOwner(ctx context.Context, id uuid.UUID) (Owner, error)
	GetProperty(ctx context.Context, id uuid.UUID) (Property, error)
	GetUser(ctx context.Context, id uuid.UUID) (User, error)
	GetUserByEmail(ctx context.Context, email string) (User, error)
	ListAgents(ctx context.Context) ([]Agent, error)
	ListOwners(ctx context.Context) ([]Owner, error)
	ListProperties(ctx context.Context, limit int32, offset int32) ([]Property, error)
	UpdateAgent(ctx context.Context, iD uuid.UUID, phoneNumber string, nationalID string, kraPin string) error
	UpdateOwner(ctx context.Context, iD uuid.UUID, phoneNumber string, nationalID string) error
	UpdateProperty(ctx context.Context, arg UpdatePropertyParams) error
	UpdateUser(ctx context.Context, iD uuid.UUID, username string, fullName string, email string, hashedPassword string) error
}

var _ Querier = (*Queries)(nil)
