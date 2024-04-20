package tests

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ngenohkevin/ark-realtors/internal/store"
	"log"
	"os"
	"testing"
)

var testStore store.Store

func TestMain(m *testing.M) {

	config := os.Getenv("DB_URL")
	if config == "" {
		log.Fatal("DB_URL environment variable is not set")
	}

	//config, err := utils.LoadConfig("../.")
	//if err != nil {
	//	log.Fatalf("cannot load config: %v", err)
	//}

	connPool, err := pgxpool.New(context.Background(), config)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testStore = store.NewStore(connPool)
	os.Exit(m.Run())
}
