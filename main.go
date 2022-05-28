package main

import (
	"database/sql"
	"fmt"
	"log"

	db "github.com/codernirmalnp/golang/db/sqlc"
	"github.com/codernirmalnp/golang/util"

	"github.com/codernirmalnp/golang/api"

	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot Load Configuartion", err)
	}

	conn, err := sql.Open(config.DbDriver, config.DbSource)
	fmt.Println(conn)
	if err != nil {

		log.Fatal("Cannot Connect to the database", err)

	}
	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("Cannot Create Server", err)
	}
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Cannot Start Server", err)

	}
}
