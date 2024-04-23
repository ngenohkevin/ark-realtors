package main

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ngenohkevin/ark-realtors/api"
	"github.com/ngenohkevin/ark-realtors/internal/store"
	"github.com/ngenohkevin/ark-realtors/pkg/utils"
	"log"
	"log/slog"
)

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	connPool, err := pgxpool.New(context.Background(), config.DbUrl)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	stores := store.NewStore(connPool)

	// Start the server
	server, err := api.NewServer(config, stores)

	if err != nil {
		log.Fatalf("cannot create server: %v", err)
	}
	err = server.Start(config.ServerAddr)
	slog.Any("Server started at", config.ServerAddr)
	if err != nil {
		log.Fatalf("cannot start server: %v", err)
	}
}
