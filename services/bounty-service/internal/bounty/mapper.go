package bounty

import (
	"database/sql"

	"github.com/DauntlessDev/bug-bounty-platform/services/bounty-service/internal/db"
	"github.com/google/uuid"
)

func toDomain(dbB db.Bounty) Bounty {
	desc := ""
	if dbB.Description.Valid {
		desc = dbB.Description.String
	}

	return Bounty{
		ID:          dbB.ID.String(),
		Title:       dbB.Title,
		Description: desc,
		Points:      int(dbB.Points),
	}
}

func toDBParams(b Bounty) (db.CreateBountyParams, error) {
	uuidID, err := uuid.Parse(b.ID)
	if err != nil {
		return db.CreateBountyParams{}, err
	}

	return db.CreateBountyParams{
		ID:    uuidID,
		Title: b.Title,
		Description: sql.NullString{
			String: b.Description,
			Valid:  b.Description != "",
		},
		Points: int32(b.Points),
	}, nil
}

func ToDBUpdateParams(b Bounty) (db.UpdateBountyParams, error) {
	uuidID, err := uuid.Parse(b.ID)
	if err != nil {
		return db.UpdateBountyParams{}, err
	}

	return db.UpdateBountyParams{
		ID:    uuidID,
		Title: b.Title,
		Description: sql.NullString{
			String: b.Description,
			Valid:  b.Description != "",
		},
		Points: int32(b.Points),
	}, nil
}
