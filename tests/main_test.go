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
	var config string

	//in GitHub actions, load config from env
	config1 := os.Getenv("DB_URL")
	if config1 == "" {
		// in local development, load config from .env file using config
		config2, err := utils.LoadConfig("../.")
		if err == nil {
			config = config2.DbUrl
		} else {
			log.Fatalf("cannot load config: %v", err)
		}
	} else {
		config = config1
	}

	connPool, err := pgxpool.New(context.Background(), config)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testStore = store.NewStore(connPool)
	os.Exit(m.Run())
}
