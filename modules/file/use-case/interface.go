package fileusecase

import (
	"context"

	"github.com/aryyawijaya/go-storage-with-clean-arch/entity"
)

type UseCase interface {
	Create(ctx context.Context, dto *CreateDto, fileDto *FileDto) (*entity.File, error)
	GetPublicByName(ctx context.Context, dto *GetPublicByNameDto) (string, error)
	GetPrivateByName(ctx context.Context, dto *GetPrivateByNameDto) (string, error)
	CreateBulk(ctx context.Context, dto []*CreateDto, filesDto []*FileDto) error
	Update(ctx context.Context, dto *UpdateDto) (*entity.File, error)
	DeleteBulk(ctx context.Context, dto *DeleteBulkDto) error
}

type Repo interface {
	GetByName(ctx context.Context, name string) (*entity.File, error)
	CreateWithTx(ctx context.Context, file *entity.File, fileContent []byte) (*entity.File, error)
	CreateBulkWithTx(ctx context.Context, files []*entity.File, filesContent [][]byte) error
	GetByNames(ctx context.Context, names []string) ([]*entity.File, error)
	Update(ctx context.Context, file *entity.File) (*entity.File, error)
	DeleteBulkWithTx(ctx context.Context, names []string) error
}

type FileValidator interface {
	IsUniqueFileName(ctx context.Context, name string) error
	IsUniqueFilesName(ctx context.Context, names []string) error
}
