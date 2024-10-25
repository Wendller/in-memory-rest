package main

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/in-memory-rest/database"
	"github.com/in-memory-rest/database/repositories"
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

	fmt.Println(users)

	return nil
}
