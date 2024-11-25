package main

import (
	"github.com/gauss2302/ecomm-service/ecomm-api/server"
	storer "github.com/gauss2302/ecomm-service/ecomm-api/store"
	"log"

	"github.com/gauss2302/ecomm-service/db"
)

func main() {
	db, err := db.NewDatabase()

	if err != nil {
		log.Fatalf("error opening database: %v", err)
	}

	defer db.Close()

	log.Println("Connected to database")

	st := storer.NewMySQLStorer(db.GetDB())
	_ = server.NewServer(st)
	//hdl := handler.NewHandler(srv)

}
