package repositories

import (
	"github.com/in-memory-rest/internal/database"
)

type UserRepo interface {
	FindAll() []database.UserSchema
	FindById(id database.ID) (database.UserSchema, error)
	Insert(firstName, lastName, bio string) (database.UserSchema, error)
	Update(id database.ID, user database.UserSchema) (database.UserSchema, error)
	Delete(id database.ID) (database.UserSchema, error)
}

type UserInMemoryRepo struct {
	DB database.UserDB
}

func NewUserInMemoryRepo(db database.UserDB) *UserInMemoryRepo {
	return &UserInMemoryRepo{
		DB: db,
	}
}

func (ur *UserInMemoryRepo) FindAll() []database.UserSchema {
	return listMapValues(ur.DB)

}

func listMapValues[K comparable, V any](hashMap map[K]V) []V {
	var values []V

	for _, v := range hashMap {
		values = append(values, v)
	}

	return values
}
