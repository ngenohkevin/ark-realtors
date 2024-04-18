package tests

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ngenohkevin/ark-realtors/internal/store"
	"github.com/ngenohkevin/ark-realtors/pkg/utils"
	"log"
	"os"
	"testing"
)

var testStore store.Store

func TestMain(m *testing.M) {
	config, err := utils.LoadConfig("../.")
	if err != nil {
		log.Fatalf("cannot load config: %v", err)
	}

	connPool, err := pgxpool.New(context.Background(), config.DbUrl)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testStore = store.NewStore(connPool)
	os.Exit(m.Run())
}
