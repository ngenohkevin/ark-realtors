package users

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ngenohkevin/ark-realtors/api"
	"github.com/ngenohkevin/ark-realtors/internal/store"
	"github.com/ngenohkevin/ark-realtors/pkg/utils"
	"github.com/stretchr/testify/require"
	"log"
	"os"
	"testing"
	"time"
)

var testStore store.Store

func newTestServer(t *testing.T, store store.Store) *api.Server {
	config := utils.Config{
		TokenSymmetricKey:   utils.RandomString(32),
		AccessTokenDuration: time.Minute,
	}
	server, err := api.NewServer(config, store)
	require.NoError(t, err)

	return server
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	cwd, _ := os.Getwd()
	fmt.Println("current working directory: ", cwd)

	config, err := utils.LoadConfig(".")
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
