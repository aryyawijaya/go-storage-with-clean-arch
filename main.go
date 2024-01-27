package main

import (
	"context"
	"log"

	"github.com/aryyawijaya/go-storage-with-clean-arch/db"
	"github.com/aryyawijaya/go-storage-with-clean-arch/server"
	utilconfig "github.com/aryyawijaya/go-storage-with-clean-arch/util/config"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
)

func main() {
	config, err := utilconfig.LoadConfig(".")
	if err != nil {
		log.Fatalf("Cannot load config: %v\n", err)
	}
	connPool, err := pgxpool.New(context.Background(), config.DBSourceContainer)
	if err != nil {
		log.Fatalf("Cannot connect to database: %v\n", err)
	}

	store := db.NewStore(connPool)
	server, err := server.NewServer(store, config)
	if err != nil {
		log.Fatalf("Cannot create the server: %v\n", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatalf("Cannot start the server: %v\n", err)
	}
}
