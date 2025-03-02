package config

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

var DB *pgxpool.Pool // Use pgxpool.Pool, not sql.DB

func ConnectDatabase() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get the database URL
	databaseUrl := os.Getenv("DATABASE_URL")
	if databaseUrl == "" {
		log.Fatal("DATABASE_URL not found in environment")
	}

	// Create a connection pool
	dbpool, err := pgxpool.New(context.Background(), databaseUrl)
	if err != nil {
		log.Fatalf("Unable to create connection pool: %v\n", err)
	}

	// Assign the connection pool to the global DB variable
	DB = dbpool

	fmt.Println("Database connected successfully.")
}
