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

func (s *Service) GetBounties(ctx context.Context) ([]Bounty, error) {
	return s.repository.GetBounties(ctx)
}

func (s *Service) GetBountiesBy(ctx context.Context, bountyID string) (Bounty, error) {
	return s.repository.GetBountyByID(ctx, bountyID)
}

func (s *Service) CreateBounty(ctx context.Context, bounty *Bounty) error {
	return s.repository.CreateBounty(ctx, bounty)
}

func (s *Service) UpdateBounty(ctx context.Context, bounty *Bounty) error {
	return s.repository.UpdateBounty(ctx, bounty)
}
