package main

import (
	"log"

	"github.com/shawon325/go-crud/config"
	"github.com/shawon325/go-crud/migrations"
)

func main() {
	config.LoadEnv()
	config.ConnectDB()

	if err := migrations.Run(config.DB); err != nil {
		log.Fatal("migration failed: ", err)
	}

	log.Println("✅ Migrations completed")
}
