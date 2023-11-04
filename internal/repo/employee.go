package repo

type Employee struct {
	ID   int
	Name string
}

func (e Employee) GetID() int {
	return e.ID
}

func (e Employee) GetRole() string {
	return "normal"	
}