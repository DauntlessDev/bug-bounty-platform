package main

import (
	"log"

	"database/sql"

	"github.com/DauntlessDev/bug-bounty-platform/services/bounty-service/internal/config"
	"github.com/DauntlessDev/bug-bounty-platform/services/bounty-service/internal/server"
	_ "github.com/lib/pq" // Import the PostgreSQL driver
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	dbConn, err := sql.Open("postgres", cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer dbConn.Close()

	err = dbConn.Ping()
	if err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}
	log.Println("Successfully connected to the database!")

	s := server.NewServer(dbConn)

	err = s.Start(":" + cfg.ServerPort)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
