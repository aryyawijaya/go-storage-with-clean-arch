package filedelivery

import "github.com/aryyawijaya/go-storage-with-clean-arch/entity"

type updateRequest struct {
	Uri struct {
		Name string `uri:"name" binding:"required"`
	}
	Body struct {
		Access entity.Access `json:"access" binding:"omitempty,access"`
	}
}
