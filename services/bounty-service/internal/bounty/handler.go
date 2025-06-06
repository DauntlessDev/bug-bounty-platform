package bounty

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

type Handler struct {
	service     *Service
	middlewares []Middleware
}

type Middleware func(http.Handler) http.Handler

func NewHandler(service *Service, middlewares ...Middleware) *Handler {
	return &Handler{service: service, middlewares: middlewares}
}

func (handler *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.Handle("GET /bounties", handler.applyMiddleware(http.HandlerFunc(handler.HandleGetBounties)))
	mux.Handle("POST /bounties", handler.applyMiddleware(http.HandlerFunc(handler.HandleCreateBounty)))
	mux.Handle("GET /bounties/", handler.applyMiddleware(http.HandlerFunc(handler.HandleGetBountyByID)))
	mux.Handle("PUT /bounties/", handler.applyMiddleware(http.HandlerFunc(handler.HandleUpdateBounty)))
}

func (handler *Handler) applyMiddleware(nextHandler http.Handler) http.Handler {
	for i := len(handler.middlewares) - 1; i >= 0; i-- {
		nextHandler = handler.middlewares[i](nextHandler)
	}
	return nextHandler
}

func (handler *Handler) HandleGetBounties(writer http.ResponseWriter, request *http.Request) {
	bounties, err := handler.service.GetBounties()
	if err != nil {
		http.Error(writer, "Failed to get bounties", http.StatusInternalServerError)
		log.Printf("Error getting bounties: %v", err)
		return
	}
	writeJSON(writer, bounties, http.StatusOK)
}

func (handler *Handler) HandleGetBountyByID(writer http.ResponseWriter, request *http.Request) {
	bountyID := strings.TrimPrefix(request.URL.Path, "/bounties/")
	if bountyID == "" {
		http.Error(writer, "Bounty ID is required", http.StatusBadRequest)
		return
	}
	bounty, err := handler.service.GetBountiesBy(bountyID)
	if err != nil {
		http.Error(writer, "Failed to get bounty by ID", http.StatusInternalServerError)
		log.Printf("Error getting bounty by ID: %v", err)
		return
	}
	writeJSON(writer, bounty, http.StatusOK)
}

func (handler *Handler) HandleCreateBounty(writer http.ResponseWriter, request *http.Request) {
	var newBounty Bounty
	if err := json.NewDecoder(request.Body).Decode(&newBounty); err != nil {
		http.Error(writer, "Invalid request body", http.StatusBadRequest)
		log.Printf("Error decoding create bounty request: %v", err)
		return
	}
	if err := handler.service.CreateBounty(&newBounty); err != nil {
		http.Error(writer, "Failed to create bounty", http.StatusInternalServerError)
		log.Printf("Error creating bounty: %v", err)
		return
	}
	writeJSON(writer, newBounty, http.StatusCreated)
}

func (handler *Handler) HandleUpdateBounty(writer http.ResponseWriter, request *http.Request) {
	var updatedBounty Bounty
	if err := json.NewDecoder(request.Body).Decode(&updatedBounty); err != nil {
		http.Error(writer, "Invalid request body", http.StatusBadRequest)
		log.Printf("Error decoding update bounty request: %v", err)
		return
	}
	if err := handler.service.UpdateBounty(&updatedBounty); err != nil {
		http.Error(writer, "Failed to update bounty", http.StatusInternalServerError)
		log.Printf("Error updating bounty: %v", err)
		return
	}
	writeJSON(writer, updatedBounty, http.StatusOK)
}

func writeJSON(writer http.ResponseWriter, data any, statusCode int) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(statusCode)
	if err := json.NewEncoder(writer).Encode(data); err != nil {
		log.Printf("Failed to write JSON response: %v", err)
	}
}
