package filevalidator

import (
	"context"

	"github.com/aryyawijaya/go-storage-with-clean-arch/entity"
	fileusecase "github.com/aryyawijaya/go-storage-with-clean-arch/modules/file/use-case"
)

type fileValidator struct {
	repo fileusecase.Repo
}

func NewFileValidator(repo fileusecase.Repo) *fileValidator {
	return &fileValidator{
		repo: repo,
	}
}

func (fv *fileValidator) IsUniqueFileName(ctx context.Context, name string) error {
	file, err := fv.repo.GetByName(ctx, name)
	if err != nil && err != entity.ErrNotFound {
		return err
	}

	if file != nil {
		return entity.ErrUnique
	}

	return nil
}

func (fv *fileValidator) IsUniqueFilesName(ctx context.Context, names []string) error {
	files, err := fv.repo.GetByNames(ctx, names)
	if err != nil && err != entity.ErrNotFound {
		return err
	}

	if len(files) != 0 {
		err = entity.ErrUnique
		return err
	}

	return nil
}
