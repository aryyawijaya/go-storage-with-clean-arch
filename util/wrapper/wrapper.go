package wrapper

import (
	"net/http"

	"github.com/aryyawijaya/go-storage-with-clean-arch/entity"
	"github.com/aryyawijaya/go-storage-with-clean-arch/middleware"
	"github.com/gin-gonic/gin"
)

type Wrapper struct{}

func NewWrapper() *Wrapper {
	return &Wrapper{}
}

func (w *Wrapper) ErrResp(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func (w *Wrapper) GetStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	switch err {
	// Entity errors
	case entity.ErrUnique,
		entity.ErrNotSameLenSlice,
		entity.ErrEmptyPayload:
		return http.StatusBadRequest

	case entity.ErrNotFound:
		return http.StatusNotFound

	case entity.ErrForbiddenAccess:
		return http.StatusForbidden

	case entity.ErrAnyEntityNotCreatedToDb:
		return http.StatusInternalServerError

	// Middleware errors
	case middleware.ErrAuthorizationHeaderNotProvided,
		middleware.ErrAuthorizationHeaderFormat,
		middleware.ErrInvalidAPIKey:
		return http.StatusUnauthorized

	default:
		return http.StatusInternalServerError
	}
}
