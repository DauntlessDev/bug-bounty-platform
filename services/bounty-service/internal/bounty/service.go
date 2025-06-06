package bounty

type Repository interface {
	GetBounties() ([]Bounty, error)
	GetBountyByID(id string) (Bounty, error)
	CreateBounty(bounty *Bounty) error
	UpdateBounty(bounty *Bounty) error
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetBounties() ([]Bounty, error) {
	return s.repo.GetBounties()
}

func (s *Service) GetBountiesBy(id string) (Bounty, error) {
	return s.repo.GetBountyByID(id)
}

func (s *Service) CreateBounty(bounty *Bounty) error {
	return s.repo.CreateBounty(bounty)
}

func (s *Service) UpdateBounty(bounty *Bounty) error {
	return s.repo.UpdateBounty(bounty)
}
