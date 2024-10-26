package domain

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Biography string `json:"biography"`
}

func (u *User) UserValidFields() []string {
	return []string{"FirstName", "LastName", "Biography"}
}
