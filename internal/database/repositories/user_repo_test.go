package repositories

import (
	"github.com/google/uuid"
	"github.com/in-memory-rest/internal/database"
	"github.com/in-memory-rest/internal/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func setupUserRepo() *UserInMemoryRepo {
	db := database.NewDB()
	repo := NewRepo(db)

	return repo.UserRepo
}

func TestFindAll(t *testing.T) {
	t.Run("return all registers", func(t *testing.T) {
		userRepo := setupUserRepo()
		id := uuid.New()
		insertedUser := database.UserSchema{
			Id:        id,
			FirstName: "John",
			LastName:  "Doe",
			Biography: "Nice guy",
		}

		userRepo.DB[id] = insertedUser

		users := userRepo.FindAll()

		assert.Equal(t, id, users[0].Id)
	})

	t.Run("when there is no register", func(t *testing.T) {
		userRepo := setupUserRepo()
		users := userRepo.FindAll()

		assert.Empty(t, users)
	})
}

func TestFindById(t *testing.T) {
	t.Run("return target user", func(t *testing.T) {
		userRepo := setupUserRepo()
		id := uuid.New()
		insertedUser := database.UserSchema{
			Id:        id,
			FirstName: "John",
			LastName:  "Doe",
			Biography: "Nice guy",
		}

		userRepo.DB[id] = insertedUser

		u, _ := userRepo.FindById(id)

		assert.Equal(t, id, u.Id)
	})

	t.Run("when there is no register", func(t *testing.T) {
		userRepo := setupUserRepo()
		id := uuid.New()
		u, err := userRepo.FindById(id)

		assert.NotEqual(t, id, u.Id)
		assert.ErrorIs(t, err, domain.ErrResourceNotFound)
	})
}

func TestInsert(t *testing.T) {
	t.Run("insert new register", func(t *testing.T) {
		userRepo := setupUserRepo()
		user, _ := userRepo.Insert("John", "Doe", "Nice Guy")

		assert.NotNil(t, user.Id)
		assert.Equal(t, "John", user.FirstName)
		assert.Equal(t, "Doe", user.LastName)
		assert.Equal(t, "Nice Guy", user.Biography)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("when user already exists", func(t *testing.T) {
		userRepo := setupUserRepo()
		user, _ := userRepo.Insert("John", "Doe", "Nice Guy")

		updateParam := domain.User{
			FirstName: "Jonny",
			LastName:  "Blaze",
		}

		updated, _ := userRepo.Update(user.Id, updateParam)

		assert.Equal(t, user.Id, updated.Id)
		assert.Equal(t, updated.FirstName, updateParam.FirstName)
		assert.Equal(t, updated.LastName, updateParam.LastName)
		assert.Equal(t, updated.Biography, user.Biography)
	})

	t.Run("when user does not exists", func(t *testing.T) {
		userRepo := setupUserRepo()

		updateParam := domain.User{
			FirstName: "Jonny",
			LastName:  "Blaze",
		}

		_, err := userRepo.Update(uuid.New(), updateParam)

		assert.ErrorIs(t, err, domain.ErrResourceNotFound)
	})
}

func TestDelete(t *testing.T) {
	t.Run("when there registered user", func(t *testing.T) {
		userRepo := setupUserRepo()
		user, _ := userRepo.Insert("John", "Doe", "Nice Guy")

		d, _ := userRepo.Delete(user.Id)

		assert.Equal(t, user.Id, d.Id)
		assert.Empty(t, userRepo.DB)
	})

	t.Run("when there is no user", func(t *testing.T) {
		userRepo := setupUserRepo()

		_, err := userRepo.Delete(uuid.New())

		assert.ErrorIs(t, err, domain.ErrResourceNotFound)
	})
}
