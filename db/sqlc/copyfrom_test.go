package sqlc_test

import (
	"context"
	"testing"

	"github.com/aryyawijaya/go-storage-with-clean-arch/db/sqlc"
	"github.com/aryyawijaya/go-storage-with-clean-arch/entity"
	utilrandom "github.com/aryyawijaya/go-storage-with-clean-arch/util/random"
	"github.com/stretchr/testify/require"
)

func TestCreateFiles(t *testing.T) {
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

	nCreatedFiles, err := testQueries.CreateFiles(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, int64(n), nCreatedFiles)

	var createdNames []string
	for _, file := range arg {
		createdNames = append(createdNames, file.Name)
	}

	currFiles, err := testQueries.GetFileByNames(context.Background(), createdNames)
	require.NoError(t, err)
	require.Len(t, currFiles, n)
}
