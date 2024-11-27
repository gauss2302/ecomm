package main

import (
	"log"
	"net/http"

	"github.com/gauss2302/ecomm-service/ecomm-api/handler"
	"github.com/gauss2302/ecomm-service/ecomm-api/server"
	storer "github.com/gauss2302/ecomm-service/ecomm-api/store"

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
	srv := server.NewServer(st)
	hdl := handler.NewHandler(srv)
	r := handler.RegisterRoutes(hdl) // Get the router

	log.Printf("Starting server on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("error starting server: %v", err)
	}
}
