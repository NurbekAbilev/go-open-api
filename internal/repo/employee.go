package repo

import "time"

type Employee struct {
	ID         int       `json:"id,omitempty"`
	FirstName  string    `json:"first_name,omitempty"`
	LastName   string    `json:"last_name,omitempty"`
	PositionID string    `json:"position_id,omitempty"`
	UpdatedAt  time.Time `json:"updated_at,omitempty"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
}
