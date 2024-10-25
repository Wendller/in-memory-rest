package repositories

import (
	"github.com/in-memory-rest/internal/database"
)

type Repo struct {
	UserRepo *UserInMemoryRepo
}

func NewRepo(db *database.DB) *Repo {
	return &Repo{
		UserRepo: NewUserInMemoryRepo(db.UsersDB),
	}
}
