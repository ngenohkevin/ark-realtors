package tests

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ngenohkevin/ark-realtors/internal/store"
	"github.com/ngenohkevin/ark-realtors/pkg/utils"
	"log"
	"os"
	"testing"
)

var testStore store.Store

func TestMain(m *testing.M) {

	config1 := os.Getenv("DB_URL")

	config2, err := utils.LoadConfig("../.")
	if err != nil && !os.IsNotExist(err) {
		log.Fatalf("cannot load config: %v", err)
	}

	fmt.Printf("config1: %v\n", config1)
	fmt.Printf("config2: %v\n", config2)

	var config string

	if config1 != "" {
		config = config1
	} else if config2.DbUrl != "" {
		config = config2.DbUrl
	} else {
		log.Fatal("DB_URL environment variable and config file DB_URL are both empty")
	}
	fmt.Printf("config: %v\n", config)

	connPool, err := pgxpool.New(context.Background(), config)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testStore = store.NewStore(connPool)
	os.Exit(m.Run())
}
