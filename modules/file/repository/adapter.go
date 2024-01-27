package filerepository

import (
	"github.com/aryyawijaya/go-storage-with-clean-arch/db/sqlc"
	"github.com/aryyawijaya/go-storage-with-clean-arch/entity"
)

func sqlcToEntity(file *sqlc.File) *entity.File {
	return &entity.File{
		ID:        file.ID,
		Name:      file.Name,
		Access:    entity.Access(file.Access),
		Path:      file.Path,
		CreatedAt: file.CreatedAt,
		UpdatedAt: file.UpdatedAt,
		Ext:       file.Ext,
	}
}

func createArgUpdate(file *entity.File) *sqlc.UpdateFileParams {
	arg := &sqlc.UpdateFileParams{
		Name:   file.Name,
		Access: sqlc.NullAccess{Access: sqlc.Access(file.Access), Valid: true},
	}

	// optional fields
	if file.Access == "" {
		arg.Access = sqlc.NullAccess{}
	}

	return arg
}
