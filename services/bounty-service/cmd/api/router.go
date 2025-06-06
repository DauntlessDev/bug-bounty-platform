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
	mux.HandleFunc("/bounties", router.bountyHandler.HandleGetBounties)
	mux.HandleFunc("/bounties/", router.bountyHandler.HandleGetBountyByID) // For /bounties/{id}
	mux.HandleFunc("/bounties/create", router.bountyHandler.HandleCreateBounty)
	mux.HandleFunc("/bounties/update", router.bountyHandler.HandleUpdateBounty)
}
