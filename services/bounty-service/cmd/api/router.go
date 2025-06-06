package main

import (
	"net/http"

	"github.com/DauntlessDev/bug-bounty-platform/services/bounty-service/internal/bounty"
)

type Router struct {
	bountyHandler *bounty.Handler
}

func NewRouter(bountyService *bounty.Service) *Router {
	return &Router{
		bountyHandler: bounty.NewHandler(bountyService),
	}
}

func (router *Router) SetupRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /bounties", router.bountyHandler.HandleGetBounties)
	mux.HandleFunc("POST /bounties", router.bountyHandler.HandleCreateBounty)
	mux.HandleFunc("GET /bounties/{id}", router.bountyHandler.HandleGetBountyByID)
	mux.HandleFunc("PUT /bounties/{id}", router.bountyHandler.HandleUpdateBounty)
}
