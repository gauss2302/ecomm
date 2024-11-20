package main

import (
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

	//Store methods here
}
