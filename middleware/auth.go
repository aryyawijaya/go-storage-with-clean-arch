package middleware

import (
	"strings"

	utilconfig "github.com/aryyawijaya/go-storage-with-clean-arch/util/config"
	"github.com/gin-gonic/gin"
)

func (m *Middleware) Auth(config *utilconfig.Config) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Extract authorization header from request
		authorizationHeader := ctx.GetHeader(config.APIKey)
		if len(authorizationHeader) == 0 {
			// Client doesn't provide authorization header
			err := ErrAuthorizationHeaderNotProvided
			ctx.AbortWithStatusJSON(m.wrapper.GetStatusCode(err), m.wrapper.ErrResp(err))
			return
		}

		// Validate authorization header format
		fields := strings.Fields(authorizationHeader)
		if len(fields) != 1 {
			err := ErrAuthorizationHeaderFormat
			ctx.AbortWithStatusJSON(m.wrapper.GetStatusCode(err), m.wrapper.ErrResp(err))
			return
		}

		// Validate API key
		apiKey := fields[0]
		if apiKey != config.APIKeyValue {
			err := ErrInvalidAPIKey
			ctx.AbortWithStatusJSON(m.wrapper.GetStatusCode(err), m.wrapper.ErrResp(err))
			return
		}

		ctx.Next()
	}
}
