package main

import (
	"dompetin-api/config"
	"dompetin-api/internal/router"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db := config.NewDB()
	r := router.SetupRouter(db)

	port := os.Getenv("PORT")
	r.Run("0.0.0.0:" + port)
}
