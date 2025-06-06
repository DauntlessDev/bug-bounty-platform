package bounty

type Bounty struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Points      int    `json:"points"`
}
