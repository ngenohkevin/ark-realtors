package main

import (
	"context"
	"errors"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
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

	runDBMigrations(config.MigrationURL, config.DbUrl)

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

func runDBMigrations(migrationURL string, dbURL string) {
	migration, err := migrate.New(migrationURL, dbURL)
	if err != nil {
		log.Fatal("cannot create migration:", err)
	}
	if err = migration.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatal("failed to run migration up:", err)
	}
	log.Println("migration successful")
}
