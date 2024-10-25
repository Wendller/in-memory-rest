package repositories

import "github.com/in-memory-rest/database"

type Repo struct {
	UserRepo *UserInMemoryRepo
}

func NewRepo(db *database.DB) *Repo {
	return &Repo{
		UserRepo: NewUserInMemoryRepo(db.UsersDB),
	}
}
