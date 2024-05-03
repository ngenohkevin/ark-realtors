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

var _ store.Store

func TestMain(m *testing.M) {
	var config string

	config1 := os.Getenv("DB_URL")
	if config1 == "" {
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

	_ = store.NewStore(connPool)
	os.Exit(m.Run())
}
