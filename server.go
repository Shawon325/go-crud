package main

import (
	"log"
	"net/http"
	"os"

	"github.com/shawon325/go-crud/config"
	"github.com/shawon325/go-crud/routes"
)

func main() {
	config.LoadEnv()
	config.ConnectDB()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	handler := routes.Routes()

	log.Println("🚀 Server running on :" + port)

	if err := http.ListenAndServe(":"+port, handler); err != nil {
		log.Fatal(err)
	}
}
