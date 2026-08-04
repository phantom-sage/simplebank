package main

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/phantom-sage/simplebank/api"
	db "github.com/phantom-sage/simplebank/db/sqlc"
	"github.com/phantom-sage/simplebank/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatalln("can not load configuration file", err)
	}
	connPool, err := pgxpool.New(context.Background(), config.DBConnUrl)
	if err != nil {
		log.Fatalln("can not connect to db", err)
	}

	store := db.NewStore(connPool)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatalln("cannot create server", err)
	}
	if err = server.Start(config.ServerAddress); err != nil {
		log.Fatalln("cannot start the HTTP server", err)
	}
}
