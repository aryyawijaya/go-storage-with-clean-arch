package sqlc_test

import (
	"context"
	"testing"

	"github.com/aryyawijaya/go-storage-with-clean-arch/db/sqlc"
	"github.com/aryyawijaya/go-storage-with-clean-arch/entity"
	utilrandom "github.com/aryyawijaya/go-storage-with-clean-arch/util/random"
	"github.com/stretchr/testify/require"
)

func createRandomFile(t *testing.T) *sqlc.File {
	access := utilrandom.RandomEnum[entity.Access](entity.AllAccessValues())
	arg := &sqlc.CreateFileParams{
		Name:   utilrandom.RandomString(5),
		Access: sqlc.Access(access),
		Path:   utilrandom.RandomPath(),
		Ext:    utilrandom.RandomFileExt(),
	}

	file, err := testQueries.CreateFile(context.Background(), arg)
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

func TestCreateFile(t *testing.T) {
	createRandomFile(t)
}

func TestGetFileByNames(t *testing.T) {
	n := 2
	var createdFiles []*sqlc.File
	for i := 0; i < n; i++ {
		file := createRandomFile(t)
		createdFiles = append(createdFiles, file)
	}

	var createdNames []string
	for _, file := range createdFiles {
		createdNames = append(createdNames, file.Name)
	}

	currFiles, err := testQueries.GetFileByNames(context.Background(), createdNames)
	require.NoError(t, err)
	require.NotEmpty(t, currFiles)

	require.Len(t, currFiles, n)
	require.Equal(t, createdFiles, currFiles)
}

func TestUpdateFile(t *testing.T) {
	t.Run("udpate all fields are optional", func(t *testing.T) {
		currFile := createRandomFile(t)

		arg := &sqlc.UpdateFileParams{
			Name: currFile.Name,
		}

		updatedFile, err := testQueries.UpdateFile(context.Background(), arg)
		require.NoError(t, err)
		require.NotEmpty(t, updatedFile)

		require.Equal(t, currFile, updatedFile)
	})

	t.Run("udpate with given value", func(t *testing.T) {
		currFile := createRandomFile(t)

		updateAccess := sqlc.Access(entity.AccessPRIVATE)
		if currFile.Access == updateAccess {
			updateAccess = sqlc.Access(entity.AccessPUBLIC)
		}

		arg := &sqlc.UpdateFileParams{
			Name:   currFile.Name,
			Access: sqlc.NullAccess{Access: updateAccess, Valid: true},
		}

		updatedFile, err := testQueries.UpdateFile(context.Background(), arg)
		require.NoError(t, err)
		require.NotEmpty(t, updatedFile)

		require.NotEqual(t, currFile, updatedFile)
		require.Equal(t, updateAccess, updatedFile.Access)
	})
}

func TestDeleteFilesByNames(t *testing.T) {
	n := 2
	var createdFiles []*sqlc.File
	for i := 0; i < n; i++ {
		file := createRandomFile(t)
		createdFiles = append(createdFiles, file)
	}

	var createdNames []string
	for _, file := range createdFiles {
		createdNames = append(createdNames, file.Name)
	}

	files, err := testQueries.DeleteFilesByNames(context.Background(), createdNames)
	require.NoError(t, err)
	require.NotEmpty(t, files)
	require.Len(t, files, n)

	currFiles, err := testQueries.GetFileByNames(context.Background(), createdNames)
	require.NoError(t, err)
	require.Empty(t, currFiles)
	require.Len(t, currFiles, 0)
}
