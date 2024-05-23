package main

import (
	"database/sql"
	"github.com/aifuxi/go-api/api"
	"github.com/aifuxi/go-api/db/sqlc"
	_ "github.com/lib/pq"
	"log"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://root:123456@localhost:5432/fuxiaochen_go_api?sslmode=disable"
	serverAddress = "127.0.0.1:8080"
)

func main() {
	db, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatalln("cannot connect to database:", err)
	}

	store := sqlc.New(db)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatalln("cannot start server:", err)
	}
}
