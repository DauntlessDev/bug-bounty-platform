package bounty

import (
	"database/sql"

	"github.com/DauntlessDev/bug-bounty-platform/services/bounty-service/internal/db"
	"github.com/google/uuid"
)

func toDomain(databaseBounty db.Bounty) Bounty {
	description := ""
	if databaseBounty.Description.Valid {
		description = databaseBounty.Description.String
	}

	return Bounty{
		ID:          databaseBounty.ID.String(),
		Title:       databaseBounty.Title,
		Description: description,
		Points:      int(databaseBounty.Points),
	}
}

func toDBParams(bounty Bounty) (db.CreateBountyParams, error) {
	uuidID, err := uuid.Parse(bounty.ID)
	if err != nil {
		return db.CreateBountyParams{}, err
	}

	return db.CreateBountyParams{
		ID:    uuidID,
		Title: bounty.Title,
		Description: sql.NullString{
			String: bounty.Description,
			Valid:  bounty.Description != "",
		},
		Points: int32(bounty.Points),
	}, nil
}

func toDBUpdateParams(bounty Bounty) (db.UpdateBountyParams, error) {
	uuidID, err := uuid.Parse(bounty.ID)
	if err != nil {
		return db.UpdateBountyParams{}, err
	}

	return db.UpdateBountyParams{
		ID:    uuidID,
		Title: bounty.Title,
		Description: sql.NullString{
			String: bounty.Description,
			Valid:  bounty.Description != "",
		},
		Points: int32(bounty.Points),
	}, nil
}
