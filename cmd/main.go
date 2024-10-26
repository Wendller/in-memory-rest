package main

import (
	"fmt"
	"github.com/fatih/structs"
	"github.com/google/uuid"
	"github.com/in-memory-rest/internal/database"
	"github.com/in-memory-rest/internal/database/repositories"
	"github.com/in-memory-rest/pkg"
)

func main() {
	fmt.Println("Hello World")
	run()
}

func run() error {
	db := database.NewDB()
	repo := repositories.NewRepo(db)

	repo.UserRepo.DB[uuid.New()] = database.UserSchema{
		Id:        uuid.New(),
		FirstName: "Wend",
		LastName:  "Ten",
		Biography: "Best",
	}

	users := repo.UserRepo.FindAll()

	user := users[0]
	userMap := structs.Map(user)
	changeset := pkg.NewChangesetValidator()

	fields := []string{"FirstName", "LastName", "Bio"}

	fmt.Println(changeset.ValidateRequired(userMap, fields))

	return nil
}
