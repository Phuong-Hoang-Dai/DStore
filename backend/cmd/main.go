package main

import (
	"log"

	config "github.com/Phuong-Hoang-Dai/DStore/configs"
	"github.com/Phuong-Hoang-Dai/DStore/db"
	"github.com/Phuong-Hoang-Dai/DStore/handler"
)

func main() {
	config.LoadConfig()
	db, err := db.SetupDB()

	if err != nil {
		log.Fatal("Error connecting database")
	}

	handler.SetupHttp(db)
}
