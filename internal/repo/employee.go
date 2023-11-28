package repo

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Employee struct {
	ID         string             `json:"id,omitempty"`
	FirstName  string             `json:"first_name,omitempty"`
	LastName   string             `json:"last_name,omitempty"`
	PositionID string             `json:"position_id,omitempty"`
	UpdatedAt  pgtype.Timestamptz `json:"updated_at,omitempty"`
	CreatedAt  pgtype.Timestamptz `json:"created_at,omitempty"`
}
