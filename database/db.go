package database

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

var DB *pgxpool.Pool

// InitDB initializes the PostgreSQL database
func InitDB() {
	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		log.Fatal("DATABASE_URL is not set in environment variables")
	}

	var err error
	DB, err = pgxpool.Connect(context.Background(), connStr)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	// Ensure the "watch_later" table exists
	createTableQuery := `
	CREATE TABLE IF NOT EXISTS watch_later (
		id SERIAL PRIMARY KEY,
		user_id INT NOT NULL,
		movie_id INT NOT NULL,
		title TEXT NOT NULL,
		category TEXT NOT NULL
	);`
	_, err = DB.Exec(context.Background(), createTableQuery)
	if err != nil {
		log.Fatalf("Failed to create watch_later table: %v\n", err)
	}

	log.Println("Connected to PostgreSQL and ensured tables exist.")
}
