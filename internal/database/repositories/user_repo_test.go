package repositories

import (
	"github.com/google/uuid"
	"github.com/in-memory-rest/internal/database"
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
