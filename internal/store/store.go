package store

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	db "github.com/ngenohkevin/ark-realtors/db/sqlc"
)

//Store defines all functions to execute db queries and transactions

type Store interface {
	db.Querier
	CreateUser(ctx context.Context, arg db.CreateUserParams) (db.User, error)
}

// SQLStore provides all functions to execute db queries and transactions
type SQLStore struct {
	connPool *pgxpool.Pool
	*db.Queries
}

// NewStore creates a new store
func NewStore(connPool *pgxpool.Pool) Store {
	return &SQLStore{
		connPool: connPool,
		Queries:  db.New(connPool),
	}
}
