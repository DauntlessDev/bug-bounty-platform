package bounty

import (
	"encoding/json"
	"log"
	"net/http"
)

type Handler struct {
	service     *Service
	middlewares []Middleware
}

type Middleware func(http.Handler) http.Handler

func NewHandler(service *Service, middlewares ...Middleware) *Handler {
	return &Handler{service: service, middlewares: middlewares}
}

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.Handle("GET /bounties", h.applyMiddleware(http.HandlerFunc(h.HandleGetBounties)))
	mux.Handle("POST /bounties", h.applyMiddleware(http.HandlerFunc(h.HandleCreateBounty)))
	mux.Handle("GET /bounties/{id}", h.applyMiddleware(http.HandlerFunc(h.HandleGetBountyByID)))
	mux.Handle("PATCH /bounties/{id}", h.applyMiddleware(http.HandlerFunc(h.HandleUpdateBounty)))
}

func (h *Handler) applyMiddleware(nextHandler http.Handler) http.Handler {
	for i := len(h.middlewares) - 1; i >= 0; i-- {
		nextHandler = h.middlewares[i](nextHandler)
	}
	return nextHandler
}

func (h *Handler) HandleGetBounties(writer http.ResponseWriter, request *http.Request) {
	bounties, err := h.service.GetBounties(request.Context())
	if err != nil {
		http.Error(writer, "Failed to get bounties", http.StatusInternalServerError)
		log.Printf("Error getting bounties: %v", err)
		return
	}
	writeJSON(writer, bounties, http.StatusOK)
}

func (h *Handler) HandleGetBountyByID(writer http.ResponseWriter, request *http.Request) {
	bountyID := request.PathValue("id")
	if bountyID == "" {
		http.Error(writer, "Bounty ID is required", http.StatusBadRequest)
		return
	}
	bounty, err := h.service.GetBountiesBy(request.Context(), bountyID)
	if err != nil {
		http.Error(writer, "Failed to get bounty by ID", http.StatusInternalServerError)
		log.Printf("Error getting bounty by ID: %v", err)
		return
	}
	writeJSON(writer, bounty, http.StatusOK)
}

func (h *Handler) HandleCreateBounty(writer http.ResponseWriter, request *http.Request) {
	var newBounty Bounty
	if err := json.NewDecoder(request.Body).Decode(&newBounty); err != nil {
		http.Error(writer, "Invalid request body", http.StatusBadRequest)
		log.Printf("Error decoding create bounty request: %v", err)
		return
	}
	if err := h.service.CreateBounty(request.Context(), &newBounty); err != nil {
		http.Error(writer, "Failed to create bounty", http.StatusInternalServerError)
		log.Printf("Error creating bounty: %v", err)
		return
	}
	writeJSON(writer, newBounty, http.StatusCreated)
}

func (h *Handler) HandleUpdateBounty(writer http.ResponseWriter, request *http.Request) {
	bountyID := request.PathValue("id")
	if bountyID == "" {
		http.Error(writer, "Bounty ID is required", http.StatusBadRequest)
		return
	}

	var updatedBounty Bounty
	if err := json.NewDecoder(request.Body).Decode(&updatedBounty); err != nil {
		http.Error(writer, "Invalid request body", http.StatusBadRequest)
		log.Printf("Error decoding update bounty request: %v", err)
		return
	}
	updatedBounty.ID = bountyID // Ensure the ID from the path is used
	if err := h.service.UpdateBounty(request.Context(), &updatedBounty); err != nil {
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
