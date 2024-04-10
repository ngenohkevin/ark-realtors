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

	url := os.Getenv("DB_URL")

	//config, err := utils.LoadConfig("..")
	//if err != nil {
	//	log.Fatalf("cannot load config: %v", err)
	//}

	connPool, err := pgxpool.New(context.Background(), url)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testStore = store.NewStore(connPool)
	os.Exit(m.Run())
}
