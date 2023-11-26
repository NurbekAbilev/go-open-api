package repo

type Employee struct {
	ID         string `json:"id,omitempty"`
	FirstName  string `json:"first_name,omitempty"`
	LastName   string `json:"last_name,omitempty"`
	PositionID int    `json:"position_id,omitempty"`
	UpdatedAt  string `json:"updated_at,omitempty"`
	CreatedAt  string `json:"created_at,omitempty"`
}

func (e Employee) GetID() string {
	return e.ID
}

func (e Employee) GetRole() string {
	return "normal"
}
