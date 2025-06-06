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
	GetBounties(context.Context) ([]db.Bounty, error)
	GetBountyByID(context.Context, uuid.UUID) (db.Bounty, error)
	CreateBounty(context.Context, db.CreateBountyParams) error
	UpdateBounty(context.Context, db.UpdateBountyParams) error
}

func NewDBRepository(queries DBQuerier) Repository {
	return &DBRepository{dbQuerier: queries}
}

func (repository *DBRepository) GetBounties() ([]Bounty, error) {
	databaseBounties, err := repository.dbQuerier.GetBounties(context.Background())
	if err != nil {
		return nil, err
	}

	bounties := make([]Bounty, len(databaseBounties))
	for i, dbItem := range databaseBounties {
		bounties[i] = toDomain(dbItem)
	}
	return bounties, nil
}

func (repository *DBRepository) GetBountyByID(bountyID string) (Bounty, error) {
	uuidID, err := uuid.Parse(bountyID)
	if err != nil {
		return Bounty{}, err
	}
	databaseBounty, err := repository.dbQuerier.GetBountyByID(context.Background(), uuidID)
	if err != nil {
		return Bounty{}, err
	}
	return toDomain(databaseBounty), nil
}

func (repository *DBRepository) CreateBounty(bounty *Bounty) error {
	params, err := toDBParams(*bounty)
	if err != nil {
		return err
	}
	return repository.dbQuerier.CreateBounty(context.Background(), params)
}

func (repository *DBRepository) UpdateBounty(bounty *Bounty) error {
	params, err := ToDBUpdateParams(*bounty)
	if err != nil {
		return err
	}
	return repository.dbQuerier.UpdateBounty(context.Background(), params)
}
