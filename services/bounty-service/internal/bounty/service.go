package bounty

type Repository interface {
	GetBounties() ([]Bounty, error)
	GetBountyByID(id string) (Bounty, error)
	CreateBounty(bounty *Bounty) error
	UpdateBounty(bounty *Bounty) error
}

type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{repository: repository}
}

func (s *Service) GetBounties() ([]Bounty, error) {
	return s.repository.GetBounties()
}

func (s *Service) GetBountiesBy(id string) (Bounty, error) {
	return s.repository.GetBountyByID(id)
}

func (s *Service) CreateBounty(bounty *Bounty) error {
	return s.repository.CreateBounty(bounty)
}

func (s *Service) UpdateBounty(bounty *Bounty) error {
	return s.repository.UpdateBounty(bounty)
}
