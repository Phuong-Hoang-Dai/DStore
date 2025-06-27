package main

import (
	"log"

	"github.com/Phuong-Hoang-Dai/DStore/db"
	"github.com/Phuong-Hoang-Dai/DStore/handler"
)

func main() {
	db, err := db.SetupDB()

	if err != nil {
		log.Fatal("Error connecting database")
	}

	handler.SetupHttp(db)
}
