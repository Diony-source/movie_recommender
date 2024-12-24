package main

import (
	"log"
	"movie_recommender/database"
	"movie_recommender/handlers"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize the database connection
	database.InitDB()

	// Start the CLI interface
	handlers.StartCLI()
}
