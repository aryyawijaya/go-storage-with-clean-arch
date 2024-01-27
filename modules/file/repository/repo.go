package filerepository

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"

	"github.com/aryyawijaya/go-storage-with-clean-arch/db"
	"github.com/aryyawijaya/go-storage-with-clean-arch/db/sqlc"
	"github.com/aryyawijaya/go-storage-with-clean-arch/entity"
	fileusecase "github.com/aryyawijaya/go-storage-with-clean-arch/modules/file/use-case"
	"github.com/jackc/pgx/v5"
)

type fileRepository struct {
	store db.Store
}

func NewFileRepository(store db.Store) fileusecase.Repo {
	return &fileRepository{
		store: store,
	}
}

func (fr *fileRepository) GetByName(ctx context.Context, name string) (*entity.File, error) {
	currFile, err := fr.store.GetFileByName(ctx, name)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, entity.ErrNotFound
		}
		return nil, err
	}

	entityFile := sqlcToEntity(currFile)

	return entityFile, nil
}

func (fr *fileRepository) CreateWithTx(ctx context.Context, file *entity.File, fileContent []byte) (*entity.File, error) {
	arg := &sqlc.CreateFileParams{
		Name:   file.Name,
		Access: sqlc.Access(file.Access),
		Path:   file.Path,
		Ext:    file.Ext,
	}

	saveFile := func() error {
		// Convert the file content byte slice to an io.Reader interface
		src := bytes.NewReader(fileContent)

		// Create path
		storePath := fmt.Sprintf("%s%s", fileusecase.PrefixPath, file.Path)
		os.MkdirAll(storePath, os.ModePerm)

		// Create a new file on the local disk
		fullPath := fmt.Sprintf("%s%s%s%s", fileusecase.PrefixPath, file.Path, file.Name, file.Ext)
		dst, err := os.Create(fullPath)
		if err != nil {
			return err
		}
		defer dst.Close()

		// Copy the uploaded file to the local disk
		_, err = io.Copy(dst, src)
		if err != nil {
			return err
		}

		return nil
	}

	result, err := fr.store.CreateFileWithTx(ctx, arg, saveFile)
	if err != nil {
		return nil, err
	}

	entityFile := sqlcToEntity(result)

	return entityFile, nil
}

func (fr *fileRepository) GetByNames(ctx context.Context, names []string) (files []*entity.File, err error) {
	currFiles, err := fr.store.GetFileByNames(ctx, names)
	if err != nil {
		if err == pgx.ErrNoRows {
			err = entity.ErrNotFound
			return
		}
		return
	}

	for _, file := range currFiles {
		entityFile := sqlcToEntity(file)

		files = append(files, entityFile)
	}

	return
}

func (fr *fileRepository) CreateBulkWithTx(ctx context.Context, files []*entity.File, filesContent [][]byte) (err error) {
	var arg []*sqlc.CreateFilesParams

	for _, file := range files {
		currFile := &sqlc.CreateFilesParams{
			Name:   file.Name,
			Access: sqlc.Access(file.Access),
			Path:   file.Path,
			Ext:    file.Ext,
		}

		arg = append(arg, currFile)
	}

	saveFiles := func() error {
		for i := 0; i < len(filesContent); i++ {
			fileContent := filesContent[i]
			file := files[i]

			// Convert the file content byte slice to an io.Reader interface
			src := bytes.NewReader(fileContent)

			// Create path
			storePath := fmt.Sprintf("%s%s", fileusecase.PrefixPath, file.Path)
			os.MkdirAll(storePath, os.ModePerm)

			// Create a new file on the local disk
			fullPath := fmt.Sprintf("%s%s%s%s", fileusecase.PrefixPath, file.Path, file.Name, file.Ext)
			dst, err := os.Create(fullPath)
			if err != nil {
				return err
			}
			defer dst.Close()

			// Copy the uploaded file to the local disk
			_, err = io.Copy(dst, src)
			if err != nil {
				return err
			}
		}
		return nil
	}

	err = fr.store.CreateFilesWithTx(ctx, arg, saveFiles)
	if err != nil {
		return
	}

	return
}

func (fr *fileRepository) Update(ctx context.Context, file *entity.File) (*entity.File, error) {
	arg := createArgUpdate(file)

	updatedFile, err := fr.store.UpdateFile(ctx, arg)
	if err != nil {
		return nil, err
	}

	entityFile := sqlcToEntity(updatedFile)

	return entityFile, nil
}

func (fr *fileRepository) DeleteBulkWithTx(ctx context.Context, names []string) error {
	deleteFiles := func(files []*sqlc.File) error {
		for _, file := range files {
			currFullPath := fmt.Sprintf("%s%s%s%s", fileusecase.PrefixPath, file.Path, file.Name, file.Ext)

			err := os.Remove(currFullPath)
			if err != nil {
				return err
			}
		}

		return nil
	}

	return fr.store.DeleteFilesWithTx(ctx, names, deleteFiles)
}
