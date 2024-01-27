package server

import (
	"github.com/aryyawijaya/go-storage-with-clean-arch/db"
	filedelivery "github.com/aryyawijaya/go-storage-with-clean-arch/modules/file/delivery"
	filerepository "github.com/aryyawijaya/go-storage-with-clean-arch/modules/file/repository"
	fileusecase "github.com/aryyawijaya/go-storage-with-clean-arch/modules/file/use-case"
	filevalidator "github.com/aryyawijaya/go-storage-with-clean-arch/modules/file/validator"
	utilconfig "github.com/aryyawijaya/go-storage-with-clean-arch/util/config"
	"github.com/aryyawijaya/go-storage-with-clean-arch/util/wrapper"
	"github.com/gin-gonic/gin"
)

type Server struct {
	Router *gin.Engine
	store  db.Store
	Config *utilconfig.Config
}

func NewServer(store db.Store, config *utilconfig.Config) (*Server, error) {
	server := &Server{
		store:  store,
		Config: config,
	}

	router := gin.Default()

	// other dependencies
	wrapper := wrapper.NewWrapper()

	// file
	fRepo := filerepository.NewFileRepository(server.store)
	fValidator := filevalidator.NewFileValidator(fRepo)
	fUseCase := fileusecase.NewFileUseCase(fRepo, fValidator)
	filedelivery.NewFileHanlder(router, fUseCase, wrapper, server.Config)

	server.Router = router

	return server, nil
}

func (s *Server) Start(address string) error {
	return s.Router.Run(address)
}
