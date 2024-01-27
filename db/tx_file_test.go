package db_test

import (
	"context"
	"errors"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/require"
)

func TestCreateFileWithTx(t *testing.T) {
	t.Run("rollback successfully", func(t *testing.T) {
		access := utilrandom.RandomEnum[entity.Access](entity.AllAccessValues())
		arg := &sqlc.CreateFileParams{
			Name:   utilrandom.RandomString(5),
			Access: sqlc.Access(access),
			Path:   utilrandom.RandomPath(),
			Ext:    utilrandom.RandomFileExt(),
		}

		errRollback := errors.New("should be rollback")
		saveFile := func() error {
			return errors.New("should be rollback")
		}

		file, err := testStore.CreateFileWithTx(context.Background(), arg, saveFile)
		require.Error(t, err)
		require.EqualError(t, err, errRollback.Error())
		require.Empty(t, file)

		file, err = testStore.GetFileByName(context.Background(), arg.Name)
		require.EqualError(t, err, pgx.ErrNoRows.Error())
		require.Empty(t, file)
	})
}

func TestCreateFilesWithTx(t *testing.T) {
	t.Run("rollback successfully", func(t *testing.T) {
		var arg []*sqlc.CreateFilesParams
		n := 2

		for i := 0; i < n; i++ {
			access := utilrandom.RandomEnum[entity.Access](entity.AllAccessValues())
			currFile := &sqlc.CreateFilesParams{
				Name:   utilrandom.RandomString(5),
				Access: sqlc.Access(access),
				Path:   utilrandom.RandomPath(),
				Ext:    utilrandom.RandomFileExt(),
			}

			arg = append(arg, currFile)
		}

		errRollback := errors.New("should be rollback")
		saveFiles := func() error {
			return errRollback
		}
		err := testStore.CreateFilesWithTx(context.Background(), arg, saveFiles)
		require.Error(t, err)
		require.EqualError(t, err, errRollback.Error())

		var names []string
		for _, file := range arg {
			names = append(names, file.Name)
		}
		files, err := testStore.GetFileByNames(context.Background(), names)
		require.NoError(t, err)
		require.Empty(t, files)
		require.Len(t, files, 0)
	})
}

func createRandomFile(t *testing.T) *sqlc.File {
	access := utilrandom.RandomEnum[entity.Access](entity.AllAccessValues())
	arg := &sqlc.CreateFileParams{
		Name:   utilrandom.RandomString(5),
		Access: sqlc.Access(access),
		Path:   utilrandom.RandomPath(),
		Ext:    utilrandom.RandomFileExt(),
	}

	file, err := testStore.CreateFile(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, file)

	require.Equal(t, arg.Name, file.Name)
	require.Equal(t, arg.Access, file.Access)
	require.Equal(t, arg.Path, file.Path)

	require.NotZero(t, file.ID)
	require.NotZero(t, file.CreatedAt)
	require.NotZero(t, file.UpdatedAt)

	return file
}

func TestDeleteFilesWithTx(t *testing.T) {
	t.Run("rollback successfully", func(t *testing.T) {
		var names []string
		n := 2

		for i := 0; i < n; i++ {
			currFile := createRandomFile(t)

			names = append(names, currFile.Name)
		}

		errRollback := errors.New("should be rollback")
		deleteFiles := func(files []*sqlc.File) error {
			return errRollback
		}
		err := testStore.DeleteFilesWithTx(context.Background(), names, deleteFiles)
		require.Error(t, err)
		require.EqualError(t, err, errRollback.Error())

		files, err := testStore.GetFileByNames(context.Background(), names)
		require.NoError(t, err)
		require.NotEmpty(t, files)
		require.Len(t, files, n)
	})
}
