package repo

import "time"

type User struct {
	ID        string    `json:"id"`
	Login     string    `json:"login"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

func (e User) GetID() int {
	return 1
	// return e.ID
}

func (e User) GetRole() string {
	return "normal"
}
