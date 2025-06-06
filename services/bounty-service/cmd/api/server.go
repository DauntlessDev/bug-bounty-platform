package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/DauntlessDev/bug-bounty-platform/services/bounty-service/internal/bounty"
	"github.com/DauntlessDev/bug-bounty-platform/services/bounty-service/internal/db"
)

type Server struct {
	router *Router
}

func NewServer(dbConn *sql.DB) *Server {
	queries := db.New(dbConn)
	bountyRepo := bounty.NewDBRepository(queries) // Use the new DBRepository
	bountyService := bounty.NewService(bountyRepo)
	router := NewRouter(bountyService)

	return &Server{
		router: router,
	}
}

func (server *Server) Start(address string) error {
	mux := http.NewServeMux()
	server.router.SetupRoutes(mux)

	log.Printf("Starting server on %s", address)
	return http.ListenAndServe(address, mux)
}
