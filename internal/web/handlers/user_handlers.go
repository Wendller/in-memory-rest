package handlers

import (
	"github.com/in-memory-rest/internal/database/repositories"
	"net/http"
)

type UserHandler struct {
	UserRepo repositories.UserRepo
}

func NewUserHandler(userRepo repositories.UserRepo) *UserHandler {
	return &UserHandler{
		UserRepo: userRepo,
	}
}

func (uh *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	allUsers := uh.UserRepo.FindAll()

	sendJSON(w, Response{Data: allUsers}, http.StatusOK)
}
