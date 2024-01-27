package middleware_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aryyawijaya/go-storage-with-clean-arch/middleware"
	utilrandom "github.com/aryyawijaya/go-storage-with-clean-arch/util/random"
	testhelper "github.com/aryyawijaya/go-storage-with-clean-arch/util/test-helper"
	"github.com/aryyawijaya/go-storage-with-clean-arch/util/wrapper"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func TestAuth(t *testing.T) {
	// Other dependencies
	wrapper := wrapper.NewWrapper()

	// Create test server
	server := testhelper.NewTestServer(t)

	t.Run("success", func(t *testing.T) {
		recorder := httptest.NewRecorder()

		request, err := http.NewRequest(http.MethodGet, testhelper.AuthMidTestPath, nil)
		require.NoError(t, err)

		// set authorization header
		request.Header.Set(server.Config.APIKey, server.Config.APIKeyValue)

		server.Router.ServeHTTP(recorder, request)

		// validate response
		require.Equal(t, http.StatusOK, recorder.Code)
	})

	t.Run("authorization not provided", func(t *testing.T) {
		recorder := httptest.NewRecorder()

		request, err := http.NewRequest(http.MethodGet, testhelper.AuthMidTestPath, nil)
		require.NoError(t, err)

		server.Router.ServeHTTP(recorder, request)

		// validate response
		require.Equal(t, http.StatusUnauthorized, recorder.Code)
		expectedResp := wrapper.ErrResp(middleware.ErrAuthorizationHeaderNotProvided)
		testhelper.RequireBodyMatchEntity[gin.H](t, &expectedResp, recorder.Body)
	})

	t.Run("invalid format", func(t *testing.T) {
		recorder := httptest.NewRecorder()

		request, err := http.NewRequest(http.MethodGet, testhelper.AuthMidTestPath, nil)
		require.NoError(t, err)

		mockAPIKeyValue := fmt.Sprintf("%s %s", utilrandom.RandomString(5), utilrandom.RandomString(5))
		request.Header.Set(server.Config.APIKey, mockAPIKeyValue)

		server.Router.ServeHTTP(recorder, request)

		// validate response
		require.Equal(t, http.StatusUnauthorized, recorder.Code)
		expectedResp := wrapper.ErrResp(middleware.ErrAuthorizationHeaderFormat)
		testhelper.RequireBodyMatchEntity[gin.H](t, &expectedResp, recorder.Body)
	})

	t.Run("invalid API key", func(t *testing.T) {
		recorder := httptest.NewRecorder()

		request, err := http.NewRequest(http.MethodGet, testhelper.AuthMidTestPath, nil)
		require.NoError(t, err)

		// set authorization header
		request.Header.Set(server.Config.APIKey, utilrandom.RandomString(8))

		server.Router.ServeHTTP(recorder, request)

		// validate response
		require.Equal(t, http.StatusUnauthorized, recorder.Code)
		expectedResp := wrapper.ErrResp(middleware.ErrInvalidAPIKey)
		testhelper.RequireBodyMatchEntity[gin.H](t, &expectedResp, recorder.Body)
	})
}
