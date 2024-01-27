package fileusecase

import "github.com/aryyawijaya/go-storage-with-clean-arch/entity"

type CreateDto struct {
	Name      string        `form:"name" binding:"required" validate:"required"`
	Access    entity.Access `form:"access" binding:"required,access" validate:"required,access"`
	Tribe     string        `form:"tribe" binding:"required,dirpath" validate:"required,dirpath"`
	Service   string        `form:"service" binding:"required,dirpath" validate:"required,dirpath"`
	Module    string        `form:"module" binding:"required,dirpath" validate:"required,dirpath"`
	SubFolder string        `form:"subFolder" binding:"omitempty,dirpath" validate:"omitempty,dirpath"`
}

type GetPublicByNameDto struct {
	Name string `uri:"name" binding:"required"`
}

type GetPrivateByNameDto struct {
	Name string `uri:"name" binding:"required"`
}

type FileDto struct {
	Content []byte
	Ext     string
}

type UpdateDto struct {
	Name   string        `uri:"name" binding:"required"`
	Access entity.Access `json:"access" binding:"required,access"`
}

type DeleteBulkDto struct {
	Names []string `json:"names" binding:"required,min=1"`
}
