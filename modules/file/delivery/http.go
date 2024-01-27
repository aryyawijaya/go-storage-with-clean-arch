package filedelivery

import (
	"github.com/aryyawijaya/go-storage-with-clean-arch/middleware"
	fileusecase "github.com/aryyawijaya/go-storage-with-clean-arch/modules/file/use-case"
	utilconfig "github.com/aryyawijaya/go-storage-with-clean-arch/util/config"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type FileHandler struct {
	fileUseCase fileusecase.UseCase
	wrapper     Wrapper
	config      *utilconfig.Config
	validate    *validator.Validate
}

func NewFileHanlder(router *gin.Engine, fileUseCase fileusecase.UseCase, wrapper Wrapper, config *utilconfig.Config) {
	handler := &FileHandler{
		fileUseCase: fileUseCase,
		wrapper:     wrapper,
		config:      config,
	}

	// custom request validator
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("access", ValidAccess)
	}

	// Create validator to validate struct (outside request)
	validate := validator.New()
	validate.RegisterValidation("access", ValidAccess)

	handler.validate = validate

	// middleware
	mid := middleware.NewMiddleware(wrapper)

	// router group
	authRouter := router.Group("/").Use(mid.Auth(config))

	authRouter.POST("/files", handler.Create)
	router.GET("/files/public/:name", handler.GetPublicByName)
	authRouter.GET("/files/private/:name", handler.GetPrivateByName)
	authRouter.POST("/files/bulk", handler.CreateBulk)
	authRouter.PATCH("/files/:name", handler.Update)
	authRouter.DELETE("/files", handler.DeleteBulk)
}
