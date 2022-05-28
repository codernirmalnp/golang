package api

import (
	"os"
	"testing"
	"time"

	db "github.com/codernirmalnp/golang/db/sqlc"
	"github.com/codernirmalnp/golang/util"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func NewTestServer(t *testing.T, store db.Store) *Server {
	config := util.Config{
		TokenKey:      util.RandomString(32),
		TokenDuration: time.Minute.String(),
	}
	server, err := NewServer(config, store)
	require.NoError(t, err)
	return server
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}
