package server

import (
	"net/http"

	"github.com/DauntlessDev/bug-bounty-platform/services/bounty-service/internal/bounty"
	"github.com/DauntlessDev/bug-bounty-platform/services/bounty-service/pkg/middleware"
)

type Router struct {
	bountyHandler *bounty.Handler
}

func NewRouter(bountyService *bounty.Service) *Router {
	bountyHandler := bounty.NewHandler(bountyService, middleware.LoggingMiddleware)
	return &Router{bountyHandler: bountyHandler}
}

func (r *Router) SetupRoutes(mux *http.ServeMux) {
	r.bountyHandler.RegisterRoutes(mux)
}
