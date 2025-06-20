package server

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
	bountyRepo := bounty.NewDBRepository(queries)
	bountyService := bounty.NewService(bountyRepo)
	router := NewRouter(bountyService)

	return &Server{
		router: router,
	}
}

func (s *Server) Start(address string) error {
	mux := http.NewServeMux()
	s.router.SetupRoutes(mux)

	log.Printf("Starting server on %s", address)
	return http.ListenAndServe(address, mux)
}
