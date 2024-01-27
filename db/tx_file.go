package db

import (
	"context"
)

func (store *SQLStore) CreateFileWithTx(ctx context.Context, arg *sqlc.CreateFileParams, saveFile func() error) (*sqlc.File, error) {
	var result *sqlc.File

	err := store.exectTx(ctx, func(qTx *sqlc.Queries) error {
		// create file db
		createdFile, err := qTx.CreateFile(ctx, arg)
		if err != nil {
			return err
		}

		// save file
		err = saveFile()
		if err != nil {
			return err
		}

		result = createdFile

		return nil
	})

	return result, err
}

func (store *SQLStore) CreateFilesWithTx(ctx context.Context, arg []*sqlc.CreateFilesParams, saveFiles func() error) (err error) {
	err = store.exectTx(ctx, func(qTx *sqlc.Queries) (err error) {
		// create files db
		nCreatedFiles, err := qTx.CreateFiles(ctx, arg)
		if err != nil {
			return
		}
		if nCreatedFiles != int64(len(arg)) {
			err = entity.ErrAnyEntityNotCreatedToDb
			return
		}

		// save files
		err = saveFiles()
		if err != nil {
			return
		}

		return
	})

	return
}

func (store *SQLStore) DeleteFilesWithTx(ctx context.Context, names []string, deleteFiles func(files []*sqlc.File) error) error {
	err := store.exectTx(ctx, func(qTx *sqlc.Queries) error {
		// delete files db
		deletedFilesDb, err := qTx.DeleteFilesByNames(ctx, names)
		if err != nil {
			return err
		}

		// delete files
		err = deleteFiles(deletedFilesDb)
		if err != nil {
			return err
		}

		return nil
	})

	return err
}
