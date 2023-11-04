package repo

type Employee struct {
	ID         string `json:"id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	PositionID int    `json:"position_id"`
	Login      string `json:"login"`
	Password   string `json:"password"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}

func (e Employee) GetID() string {
	return e.ID
}

func (e Employee) GetRole() string {
	return "normal"
}
