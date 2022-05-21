package main

import (
	"database/sql"
	"fmt"
	"log"
	"test/api"
	db "test/db/sqlc"
	"test/db/util"

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
	server := api.NewServer(store)
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Cannot Start Server", err)

	}
}
