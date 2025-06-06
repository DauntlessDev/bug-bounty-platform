package bounty

import (
	"context"
	"database/sql"

	"github.com/DauntlessDev/bug-bounty-platform/services/bounty-service/internal/db"
	"github.com/google/uuid"
)

// DBRepository implements the Repository interface using db.Queries.
type DBRepository struct {
	queries *db.Queries
}

// NewDBRepository creates a new DBRepository.
func NewDBRepository(queries *db.Queries) Repository {
	return &DBRepository{queries: queries}
}

// GetBounties retrieves all bounties from the database.
func (r *DBRepository) GetBounties() ([]Bounty, error) {
	dbBounties, err := r.queries.GetBounties(context.Background())
	if err != nil {
		return nil, err
	}

	bounties := make([]Bounty, len(dbBounties))
	for i, dbBounty := range dbBounties {
		bounties[i] = Bounty{
			ID:          dbBounty.ID.String(),
			Title:       dbBounty.Title,
			Description: dbBounty.Description.String, // Handle sql.NullString
			Points:      int(dbBounty.Points),
		}
	}
	return bounties, nil
}

// GetBountyByID retrieves a bounty by its ID from the database.
func (r *DBRepository) GetBountyByID(id string) (Bounty, error) {
	uuidID, err := uuid.Parse(id)
	if err != nil {
		return Bounty{}, err // Or a specific error type for invalid ID
	}

	dbBounty, err := r.queries.GetBountyByID(context.Background(), uuidID)
	if err != nil {
		return Bounty{}, err
	}

	bounty := Bounty{
		ID:          dbBounty.ID.String(),
		Title:       dbBounty.Title,
		Description: dbBounty.Description.String, // Handle sql.NullString
		Points:      int(dbBounty.Points),
	}
	return bounty, nil
}

// CreateBounty creates a new bounty in the database.
func (r *DBRepository) CreateBounty(bounty *Bounty) error {
	uuidID, err := uuid.Parse(bounty.ID)
	if err != nil {
		return err // Or a specific error type for invalid ID
	}

	params := db.CreateBountyParams{
		ID:    uuidID,
		Title: bounty.Title,
		Description: sql.NullString{ // Handle sql.NullString
			String: bounty.Description,
			Valid:  bounty.Description != "",
		},
		Points: int32(bounty.Points),
	}

	return r.queries.CreateBounty(context.Background(), params)
}

// UpdateBounty updates an existing bounty in the database.
func (r *DBRepository) UpdateBounty(bounty *Bounty) error {
	uuidID, err := uuid.Parse(bounty.ID)
	if err != nil {
		return err // Or a specific error type for invalid ID
	}

	params := db.UpdateBountyParams{
		ID:    uuidID,
		Title: bounty.Title,
		Description: sql.NullString{ // Handle sql.NullString
			String: bounty.Description,
			Valid:  bounty.Description != "",
		},
		Points: int32(bounty.Points),
	}

	return r.queries.UpdateBounty(context.Background(), params)
}
