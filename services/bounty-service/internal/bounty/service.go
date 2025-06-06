package bounty

import "context"

type Repository interface {
	GetBounties(ctx context.Context) ([]Bounty, error)
	GetBountyByID(ctx context.Context, id string) (Bounty, error)
	CreateBounty(ctx context.Context, bounty *Bounty) error
	UpdateBounty(ctx context.Context, bounty *Bounty) error
}

type Service struct {
	repository Repository
}

func NewService(repo Repository) *Service {
	return &Service{repository: repo}
}

func (service *Service) GetBounties(ctx context.Context) ([]Bounty, error) {
	return service.repository.GetBounties(ctx)
}

func (service *Service) GetBountiesBy(ctx context.Context, bountyID string) (Bounty, error) {
	return service.repository.GetBountyByID(ctx, bountyID)
}

func (service *Service) CreateBounty(ctx context.Context, bounty *Bounty) error {
	return service.repository.CreateBounty(ctx, bounty)
}

func (service *Service) UpdateBounty(ctx context.Context, bounty *Bounty) error {
	return service.repository.UpdateBounty(ctx, bounty)
}
