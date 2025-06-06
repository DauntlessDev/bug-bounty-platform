package bounty

type Repository interface {
	GetBounties() ([]Bounty, error)
	GetBountyByID(id string) (Bounty, error)
	CreateBounty(bounty *Bounty) error
	UpdateBounty(bounty *Bounty) error
}
