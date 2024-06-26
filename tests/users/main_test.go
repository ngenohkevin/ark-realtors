package users

import (
	"github.com/gin-gonic/gin"
	"github.com/ngenohkevin/ark-realtors/api"
	"github.com/ngenohkevin/ark-realtors/internal/store"
	"github.com/ngenohkevin/ark-realtors/pkg/utils"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
	"time"
)

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
	os.Exit(m.Run())

}
