package bounty

import (
	"encoding/json"
	"log"
	"net/http"
)

// Handler struct holds the bounty service dependency.
type Handler struct {
	service *Service
}

// NewHandler creates a new Handler instance.
func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

// HandleGetBounties handles the request to get all bounties.
func (h *Handler) HandleGetBounties(w http.ResponseWriter, r *http.Request) {
	bounties, err := h.service.GetBounties()
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

// HandleGetBountyByID handles the request to get a bounty by its ID.
func (h *Handler) HandleGetBountyByID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/bounties/"):]
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

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(bounty); err != nil {
		http.Error(w, "Failed to marshal bounty to JSON", http.StatusInternalServerError)
		log.Printf("Error marshalling bounty: %v", err)
		return
	}
}

// HandleCreateBounty handles the request to create a new bounty.
func (h *Handler) HandleCreateBounty(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var newBounty Bounty // Assuming Bounty struct is in the same package
	if err := json.NewDecoder(r.Body).Decode(&newBounty); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		log.Printf("Error decoding create bounty request: %v", err)
		return
	}

	if err := h.service.CreateBounty(&newBounty); err != nil {
		http.Error(w, "Failed to create bounty", http.StatusInternalServerError)
		log.Printf("Error creating bounty: %v", err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(newBounty); err != nil {
		log.Printf("Error marshalling created bounty: %v", err)
	}
}

// HandleUpdateBounty handles the request to update an existing bounty.
func (h *Handler) HandleUpdateBounty(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var updatedBounty Bounty // Assuming Bounty struct is in the same package
	if err := json.NewDecoder(r.Body).Decode(&updatedBounty); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		log.Printf("Error decoding update bounty request: %v", err)
		return
	}

	if err := h.service.UpdateBounty(&updatedBounty); err != nil {
		http.Error(w, "Failed to update bounty", http.StatusInternalServerError)
		log.Printf("Error updating bounty: %v", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(updatedBounty); err != nil {
		log.Printf("Error marshalling updated bounty: %v", err)
	}
}
