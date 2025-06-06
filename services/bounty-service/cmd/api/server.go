package main

import (
	"database/sql"
	"encoding/json" // Import the json package
	"log"
	"net/http"

	"github.com/DauntlessDev/bug-bounty-platform/services/bounty-service/internal/db"
)

type Server struct {
	queries *db.Queries
}

func NewServer(dbConn *sql.DB) *Server {
	queries := db.New(dbConn)
	return &Server{
		queries: queries,
	}
}

func (server *Server) Start(address string) error {
	http.HandleFunc("/bounties", server.handleGetBounties)

	log.Printf("Starting server on %s", address)
	return http.ListenAndServe(address, nil)
}

func (server *Server) handleGetBounties(w http.ResponseWriter, r *http.Request) {
	bounties, err := server.queries.GetBounties(r.Context())
	if err != nil {
		http.Error(w, "Failed to get bounties", http.StatusInternalServerError)
		log.Printf("Error getting bounties: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(bounties); err != nil {
		http.Error(w, "Failed to marshal bounties to JSON", http.StatusInternalServerError)
		log.Printf("Error marshalling bounties: %v", err)
		return
	}
}
