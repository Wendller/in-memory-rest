package main

import (
	"fmt"
	"github.com/fatih/structs"
	"github.com/google/uuid"
	"github.com/in-memory-rest/configs"
	"github.com/in-memory-rest/internal/database"
	"github.com/in-memory-rest/internal/database/repositories"
	"github.com/in-memory-rest/pkg"
	"log/slog"
	"net/http"
)

func main() {
	cfg := configs.LoadConfig()

	if err := http.ListenAndServe("localhost:8080", cfg.Router); err != nil {
		slog.Error("application initialize error", "error", err)
		return
	}

	fmt.Println("start listening on port :8080")
	slog.Info("all systems initialized")
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
