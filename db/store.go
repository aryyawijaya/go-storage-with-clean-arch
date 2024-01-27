package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Store interface provide all function Queries and + transaction
type Store interface {
	sqlc.Querier
	CreateFileWithTx(ctx context.Context, arg *sqlc.CreateFileParams, saveFile func() error) (*sqlc.File, error)
	CreateFilesWithTx(ctx context.Context, arg []*sqlc.CreateFilesParams, saveFiles func() error) error
	DeleteFilesWithTx(ctx context.Context, names []string, deleteFiles func(files []*sqlc.File) error) error
}

// Store struct provide all function Queries and + transaction
type SQLStore struct {
	connPool *pgxpool.Pool
	*sqlc.Queries
}

func NewStore(connPool *pgxpool.Pool) Store {
	return &SQLStore{
		connPool: connPool,
		Queries:  sqlc.New(connPool),
	}
}

// executes function within a database transaction
func (store *SQLStore) exectTx(ctx context.Context, fn func(qTx *sqlc.Queries) error) error {
	tx, err := store.connPool.Begin(ctx)
	if err != nil {
		return err
	}

	qTx := sqlc.New(tx)
	err = fn(qTx)
	if err != nil {
		if rbErr := tx.Rollback(ctx); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit(ctx)
}
