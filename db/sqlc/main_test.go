package sqlc_test

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/aryyawijaya/go-storage-with-clean-arch/db/sqlc"
	utilconfig "github.com/aryyawijaya/go-storage-with-clean-arch/util/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

var testQueries *sqlc.Queries

func TestMain(m *testing.M) {
	config, err := utilconfig.LoadConfig("../../")
	if err != nil {
		log.Fatalf("Cannot load config: %v\n", err)
	}

	connPool, err := pgxpool.New(context.Background(), config.DBSourceLocal)
	if err != nil {
		log.Fatalf("Cannot connect to database: %v\n", err)
	}

	testQueries = sqlc.New(connPool)

	os.Exit(m.Run())
}
