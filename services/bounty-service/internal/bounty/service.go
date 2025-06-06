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

func NewService(repo Repository) *Service {
	return &Service{repository: repo}
}

func (service *Service) GetBounties() ([]Bounty, error) {
	return service.repository.GetBounties()
}

func (service *Service) GetBountiesBy(bountyID string) (Bounty, error) {
	return service.repository.GetBountyByID(bountyID)
}

func (service *Service) CreateBounty(bounty *Bounty) error {
	return service.repository.CreateBounty(bounty)
}

func (service *Service) UpdateBounty(bounty *Bounty) error {
	return service.repository.UpdateBounty(bounty)
}
