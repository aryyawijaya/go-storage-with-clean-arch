package fileusecase

import (
	"context"
	"fmt"

	"github.com/aryyawijaya/go-storage-with-clean-arch/entity"
)

const (
	PrefixPath = "stored-files/"
)

type fileUseCase struct {
	repo      Repo
	validator FileValidator
}

func NewFileUseCase(repo Repo, validator FileValidator) UseCase {
	return &fileUseCase{
		repo:      repo,
		validator: validator,
	}
}

func (f *fileUseCase) Create(ctx context.Context, dto *CreateDto, fileDto *FileDto) (file *entity.File, err error) {
	err = f.validator.IsUniqueFileName(ctx, dto.Name)
	if err != nil {
		return
	}

	filepath := fmt.Sprintf("%s%s%s%s", dto.Tribe, dto.Service, dto.Module, dto.SubFolder)

	createFile := &entity.File{
		Name:   dto.Name,
		Access: dto.Access,
		Path:   filepath,
		Ext:    fileDto.Ext,
	}

	file, err = f.repo.CreateWithTx(ctx, createFile, fileDto.Content)
	if err != nil {
		return
	}

	return
}

func (f *fileUseCase) GetPublicByName(ctx context.Context, dto *GetPublicByNameDto) (filepath string, err error) {
	file, err := f.repo.GetByName(ctx, dto.Name)
	if err != nil {
		return
	}

	if file.Access != entity.AccessPUBLIC {
		err = entity.ErrForbiddenAccess
		return
	}

	filepath = fmt.Sprintf("%s%s%s", file.Path, file.Name, file.Ext)

	return
}

func (f *fileUseCase) GetPrivateByName(ctx context.Context, dto *GetPrivateByNameDto) (filepath string, err error) {
	file, err := f.repo.GetByName(ctx, dto.Name)
	if err != nil {
		return
	}

	filepath = fmt.Sprintf("%s%s%s", file.Path, file.Name, file.Ext)

	return
}

func (f *fileUseCase) CreateBulk(ctx context.Context, dto []*CreateDto, filesDto []*FileDto) (err error) {
	// Get slice of name file in dto
	var names []string
	for _, file := range dto {
		names = append(names, file.Name)
	}

	err = f.validator.IsUniqueFilesName(ctx, names)
	if err != nil {
		return
	}

	// Mapping dto to slice of File & mapping filesDto to slice of []byte
	var createFiles []*entity.File
	var filesContent [][]byte
	for i := 0; i < len(dto); i++ {
		currDto := dto[i]
		currFileDto := filesDto[i]

		filepath := fmt.Sprintf("%s%s%s%s", currDto.Tribe, currDto.Service, currDto.Module, currDto.SubFolder)

		currFile := &entity.File{
			Name:   currDto.Name,
			Access: currDto.Access,
			Path:   filepath,
			Ext:    currFileDto.Ext,
		}

		createFiles = append(createFiles, currFile)
		filesContent = append(filesContent, currFileDto.Content)
	}

	err = f.repo.CreateBulkWithTx(ctx, createFiles, filesContent)
	if err != nil {
		return
	}

	return
}

func (f *fileUseCase) Update(ctx context.Context, dto *UpdateDto) (*entity.File, error) {
	file := &entity.File{
		Name:   dto.Name,
		Access: dto.Access,
	}

	updatedFile, err := f.repo.Update(ctx, file)
	if err != nil {
		return nil, err
	}

	return updatedFile, nil
}

func (f *fileUseCase) DeleteBulk(ctx context.Context, dto *DeleteBulkDto) error {
	err := f.repo.DeleteBulkWithTx(ctx, dto.Names)
	if err != nil {
		return err
	}

	return nil
}
