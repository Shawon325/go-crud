package main

import (
	"log"
	"net/http"
	"os"

	"github.com/shawon325/go-crud/config"
	"github.com/shawon325/go-crud/routes"
	"github.com/shawon325/go-crud/src/models"
)

func main() {
	config.LoadEnv()
	config.ConnectDB()

	if err := config.DB.AutoMigrate(&models.User{}); err != nil {
		log.Fatal("migration failed: ", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	handler := routes.Routes()

	log.Println("🚀 Server running on :" + port)
	log.Fatal(http.ListenAndServe(":"+port, handler))

}
