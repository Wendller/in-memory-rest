package handlers

import (
	"encoding/json"
	"errors"
	"github.com/fatih/structs"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/in-memory-rest/internal/database/repositories"
	"github.com/in-memory-rest/internal/domain"
	"github.com/in-memory-rest/pkg"
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

func (uh *UserHandler) Insert(w http.ResponseWriter, r *http.Request) {
	var body domain.User
	var changeset = pkg.NewChangesetValidator()

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		sendJSON(w,
			Response{Error: "Please provide first_name last_name and biography for the user"},
			http.StatusBadRequest,
		)
		return
	}

	changeset.ValidateRequired(structs.Map(body), body.UserValidFields())

	if !changeset.IsValid {
		sendJSON(w, Response{Error: changeset.Errors}, http.StatusBadRequest)
		return
	}

	changeset.MinStrLen("FirstName", body.FirstName, 2)
	changeset.MaxStrLen("FirstName", body.FirstName, 20)
	changeset.MinStrLen("LastName", body.LastName, 2)
	changeset.MaxStrLen("LastName", body.LastName, 20)
	changeset.MinStrLen("Biography", body.Biography, 20)
	changeset.MaxStrLen("Biography", body.Biography, 450)

	if !changeset.IsValid {
		sendJSON(w, Response{Error: changeset.Errors}, http.StatusBadRequest)
		return
	}

	u, err := uh.UserRepo.Insert(body.FirstName, body.LastName, body.Biography)
	if err != nil {
		sendJSON(w,
			Response{Error: "There was an error while saving the user to the database"},
			http.StatusInternalServerError,
		)
		return
	}

	sendJSON(w, Response{Data: u}, http.StatusOK)
}

func (uh *UserHandler) GetUserById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	parsedId, err := uuid.Parse(id)
	if err != nil {
		sendJSON(w, Response{Error: err.Error()}, http.StatusBadRequest)
		return
	}

	u, err := uh.UserRepo.FindById(parsedId)
	if err != nil {
		sendJSON(w, Response{Error: "The user with the specified ID does not exist"}, http.StatusNotFound)
		return
	}

	sendJSON(w, Response{Data: u}, http.StatusOK)
}

func (uh *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	var body domain.User
	var changeset = pkg.NewChangesetValidator()

	id := chi.URLParam(r, "id")

	parsedId, err := uuid.Parse(id)
	if err != nil {
		sendJSON(w, Response{Error: err.Error()}, http.StatusBadRequest)
		return
	}

	if err = json.NewDecoder(r.Body).Decode(&body); err != nil {
		sendJSON(w,
			Response{Error: "Please provide first_name last_name and biography for the user"},
			http.StatusBadRequest,
		)
		return
	}

	if body.FirstName != "" {
		changeset.MinStrLen("FirstName", body.FirstName, 2)
		changeset.MaxStrLen("FirstName", body.FirstName, 20)
	}

	if body.LastName != "" {
		changeset.MinStrLen("LastName", body.LastName, 2)
		changeset.MaxStrLen("LastName", body.LastName, 20)
	}

	if body.Biography != "" {
		changeset.MinStrLen("Biography", body.Biography, 20)
		changeset.MaxStrLen("Biography", body.Biography, 450)
	}

	if !changeset.IsValid {
		sendJSON(w, Response{Error: changeset.Errors}, http.StatusBadRequest)
		return
	}

	u, err := uh.UserRepo.Update(parsedId, body)
	if err != nil {
		if errors.Is(err, domain.ErrResourceNotFound) {
			sendJSON(w,
				Response{Error: "The user with the specified ID does not exist"},
				http.StatusNotFound,
			)
		} else {
			sendJSON(w,
				Response{Error: "The user information could not be modified"},
				http.StatusInternalServerError,
			)
		}
		return
	}

	sendJSON(w, Response{Data: u}, http.StatusOK)
}

func (uh *UserHandler) DeleteUserById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	parsedId, err := uuid.Parse(id)
	if err != nil {
		sendJSON(w, Response{Error: err.Error()}, http.StatusBadRequest)
		return
	}

	u, err := uh.UserRepo.Delete(parsedId)
	if err != nil {
		if errors.Is(err, domain.ErrResourceNotFound) {
			sendJSON(w,
				Response{Error: "The user with the specified ID does not exist"},
				http.StatusNotFound,
			)
		} else {
			sendJSON(w,
				Response{Error: "The user could not be removed"},
				http.StatusInternalServerError,
			)
		}
		return
	}

	sendJSON(w, Response{Data: u}, http.StatusOK)
}
