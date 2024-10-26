package repositories

import (
	"github.com/google/uuid"
	"github.com/in-memory-rest/internal"
	"github.com/in-memory-rest/internal/database"
	"github.com/in-memory-rest/internal/domain"
)

type UserRepo interface {
	FindAll() []database.UserSchema
	FindById(id database.ID) (database.UserSchema, error)
	Insert(firstName, lastName, bio string) (database.UserSchema, error)
	Update(id database.ID, user domain.User) (database.UserSchema, error)
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

func (ur *UserInMemoryRepo) Insert(firstName, lastName, biography string) (database.UserSchema, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return database.UserSchema{}, err
	}

	newUser := database.UserSchema{
		Id:        id,
		FirstName: firstName,
		LastName:  lastName,
		Biography: biography,
	}

	ur.DB[id] = newUser

	return newUser, nil
}

func (ur *UserInMemoryRepo) FindById(id database.ID) (database.UserSchema, error) {
	u, ok := ur.DB[id]
	if !ok {
		return database.UserSchema{}, internal.ErrResourceNotFound
	}

	return u, nil
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

func (ur *UserInMemoryRepo) Update(id database.ID, user domain.User) (database.UserSchema, error) {
	u, ok := ur.DB[id]
	if !ok {
		return database.UserSchema{}, internal.ErrResourceNotFound
	}

	if user.FirstName != "" {
		u.FirstName = user.FirstName
	}

	if user.LastName != "" {
		u.LastName = user.LastName
	}

	if user.Biography != "" {
		u.Biography = user.Biography
	}

	return u, nil
}

func (ur *UserInMemoryRepo) Delete(id database.ID) (database.UserSchema, error) {
	u, ok := ur.DB[id]
	if !ok {
		return database.UserSchema{}, internal.ErrResourceNotFound
	}

	delete(ur.DB, id)

	return u, nil
}
