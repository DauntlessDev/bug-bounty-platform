package bounty

import (
	"context"

	"github.com/DauntlessDev/bug-bounty-platform/services/bounty-service/internal/db"
	"github.com/google/uuid"
)

type DBRepository struct {
	q DBQuerier
}

type DBQuerier interface {
	GetBounties(context.Context) ([]db.Bounty, error)
	GetBountyByID(context.Context, uuid.UUID) (db.Bounty, error)
	CreateBounty(context.Context, db.CreateBountyParams) error
	UpdateBounty(context.Context, db.UpdateBountyParams) error
}

func NewDBRepository(queries DBQuerier) Repository {
	return &DBRepository{q: queries}
}

func (r *DBRepository) GetBounties() ([]Bounty, error) {
	dbList, err := r.q.GetBounties(context.Background())
	if err != nil {
		return nil, err
	}

	result := make([]Bounty, len(dbList))
	for i, dbItem := range dbList {
		result[i] = toDomain(dbItem)
	}
	return result, nil
}

func (r *DBRepository) GetBountyByID(id string) (Bounty, error) {
	uuidID, err := uuid.Parse(id)
	if err != nil {
		return Bounty{}, err
	}
	dbItem, err := r.q.GetBountyByID(context.Background(), uuidID)
	if err != nil {
		return Bounty{}, err
	}
	return toDomain(dbItem), nil
}

func (r *DBRepository) CreateBounty(b *Bounty) error {
	params, err := toDBParams(*b)
	if err != nil {
		return err
	}
	return r.q.CreateBounty(context.Background(), params)
}

func (r *DBRepository) UpdateBounty(b *Bounty) error {
	params, err := toDBUpdateParams(*b)
	if err != nil {
		return err
	}
	return r.q.UpdateBounty(context.Background(), params)
}
