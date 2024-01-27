package testhelper

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	mockdb "github.com/aryyawijaya/go-storage-with-clean-arch/db/mock"
	"github.com/aryyawijaya/go-storage-with-clean-arch/middleware"
	"github.com/aryyawijaya/go-storage-with-clean-arch/server"
	utilconfig "github.com/aryyawijaya/go-storage-with-clean-arch/util/config"
	utilrandom "github.com/aryyawijaya/go-storage-with-clean-arch/util/random"
	"github.com/aryyawijaya/go-storage-with-clean-arch/util/wrapper"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

const (
	AuthMidTestPath = "/test/mid/auth"
)

func NewTestServer(t *testing.T) *server.Server {
	// Create MockStore
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	store := mockdb.NewMockStore(ctrl)

	// Mock Config
	config := &utilconfig.Config{
		APIKey:      utilrandom.RandomString(5),
		APIKeyValue: utilrandom.RandomString(16),
	}

	s, err := server.NewServer(store, config)
	require.NoError(t, err)

	// dependencies
	w := wrapper.NewWrapper()
	mid := middleware.NewMiddleware(w)

	// auth middleware test route
	s.Router.GET(
		AuthMidTestPath,
		mid.Auth(s.Config),
		func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{})
		},
	)

	return s
}

func RequireBodyMatchEntity[E any](t *testing.T, entity *E, body *bytes.Buffer) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var gotEntity *E
	err = json.Unmarshal(data, &gotEntity)
	require.NoError(t, err)
	require.Equal(t, entity, gotEntity)
}
