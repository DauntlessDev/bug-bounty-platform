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

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.Handle("GET /bounties", h.applyMiddleware(http.HandlerFunc(h.HandleGetBounties)))
	mux.Handle("POST /bounties", h.applyMiddleware(http.HandlerFunc(h.HandleCreateBounty)))
	mux.Handle("GET /bounties/", h.applyMiddleware(http.HandlerFunc(h.HandleGetBountyByID)))
	mux.Handle("PUT /bounties/", h.applyMiddleware(http.HandlerFunc(h.HandleUpdateBounty)))
}

func (h *Handler) applyMiddleware(handler http.Handler) http.Handler {
	for i := len(h.middlewares) - 1; i >= 0; i-- {
		handler = h.middlewares[i](handler)
	}
	return handler
}

func (h *Handler) HandleGetBounties(w http.ResponseWriter, r *http.Request) {
	bounties, err := h.service.GetBounties()
	if err != nil {
		http.Error(w, "Failed to get bounties", http.StatusInternalServerError)
		log.Printf("Error getting bounties: %v", err)
		return
	}
	writeJSON(w, bounties, http.StatusOK)
}

func (h *Handler) HandleGetBountyByID(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/bounties/")
	if id == "" {
		http.Error(w, "Bounty ID is required", http.StatusBadRequest)
		return
	}
	bounty, err := h.service.GetBountiesBy(id)
	if err != nil {
		http.Error(w, "Failed to get bounty by ID", http.StatusInternalServerError)
		log.Printf("Error getting bounty by ID: %v", err)
		return
	}
	writeJSON(w, bounty, http.StatusOK)
}

func (h *Handler) HandleCreateBounty(w http.ResponseWriter, r *http.Request) {
	var bounty Bounty
	if err := json.NewDecoder(r.Body).Decode(&bounty); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		log.Printf("Error decoding create bounty request: %v", err)
		return
	}
	if err := h.service.CreateBounty(&bounty); err != nil {
		http.Error(w, "Failed to create bounty", http.StatusInternalServerError)
		log.Printf("Error creating bounty: %v", err)
		return
	}
	writeJSON(w, bounty, http.StatusCreated)
}

func (h *Handler) HandleUpdateBounty(w http.ResponseWriter, r *http.Request) {
	var bounty Bounty
	if err := json.NewDecoder(r.Body).Decode(&bounty); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		log.Printf("Error decoding update bounty request: %v", err)
		return
	}
	if err := h.service.UpdateBounty(&bounty); err != nil {
		http.Error(w, "Failed to update bounty", http.StatusInternalServerError)
		log.Printf("Error updating bounty: %v", err)
		return
	}
	writeJSON(w, bounty, http.StatusOK)
}

func writeJSON(w http.ResponseWriter, data any, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("Failed to write JSON response: %v", err)
	}
}
