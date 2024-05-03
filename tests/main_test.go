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

	config1 := os.Getenv("DB_URL")
	if config1 == "" {
		config2, err := utils.LoadConfig("../.")
		if err != nil {
			config = config2.DbUrl
		}
	}

	//if config1 != "" {
	//	config = config1
	//} else if config2.DbUrl != "" {

	//} else {
	//	log.Fatal("cannot load config")
	//}

	connPool, err := pgxpool.New(context.Background(), config)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testStore = store.NewStore(connPool)
	os.Exit(m.Run())
}
