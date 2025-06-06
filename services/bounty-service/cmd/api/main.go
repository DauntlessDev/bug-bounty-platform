package main

import (
	"fmt"
	"log"
	"os"

	"database/sql"

	"github.com/DauntlessDev/bug-bounty-platform/services/bounty-service/internal/server"
	_ "github.com/lib/pq" // Import the PostgreSQL driver
)

func main() {
	dataSourceName := os.Getenv("DATABASE_URL")
	if dataSourceName == "" {
		dataSourceName = "host=localhost port=5432 user=postgres password=postgres dbname=bounty_service sslmode=disable"
		log.Println("DATABASE_URL not set, using default connection string.")
	}

	dbConn, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer dbConn.Close()

	err = dbConn.Ping()
	if err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}
	fmt.Println("Successfully connected to the database!")

	server := server.NewServer(dbConn)

	err = server.Start(":8080")
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
