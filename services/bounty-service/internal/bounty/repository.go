package bounty

import (
	"context"

	"github.com/DauntlessDev/bug-bounty-platform/services/bounty-service/internal/db"
	"github.com/google/uuid"
)

type DBRepository struct {
	dbQuerier DBQuerier
}

type DBQuerier interface {
	GetBounties(ctx context.Context) ([]db.Bounty, error)
	GetBountyByID(ctx context.Context, id uuid.UUID) (db.Bounty, error)
	CreateBounty(ctx context.Context, arg db.CreateBountyParams) error
	UpdateBounty(ctx context.Context, arg db.UpdateBountyParams) error
}

func NewDBRepository(queries DBQuerier) Repository {
	return &DBRepository{dbQuerier: queries}
}

func (r *DBRepository) GetBounties(ctx context.Context) ([]Bounty, error) {
	databaseBounties, err := r.dbQuerier.GetBounties(ctx)
	if err != nil {
		return nil, err
	}

	bounties := make([]Bounty, len(databaseBounties))
	for i, dbItem := range databaseBounties {
		bounties[i] = toDomain(dbItem)
	}
	return bounties, nil
}

func (r *DBRepository) GetBountyByID(ctx context.Context, bountyID string) (Bounty, error) {
	uuidID, err := uuid.Parse(bountyID)
	if err != nil {
		return Bounty{}, err
	}
	databaseBounty, err := r.dbQuerier.GetBountyByID(ctx, uuidID)
	if err != nil {
		return Bounty{}, err
	}
	return toDomain(databaseBounty), nil
}

func (r *DBRepository) CreateBounty(ctx context.Context, bounty *Bounty) error {
	params, err := toDBParams(*bounty)
	if err != nil {
		return err
	}
	return r.dbQuerier.CreateBounty(ctx, params)
}

func (r *DBRepository) UpdateBounty(ctx context.Context, bounty *Bounty) error {
	params, err := toDBUpdateParams(*bounty)
	if err != nil {
		return err
	}
	return r.dbQuerier.UpdateBounty(ctx, params)
}
