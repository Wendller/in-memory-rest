package handlers

import (
	"encoding/json"
	"github.com/in-memory-rest/internal/database/repositories"
	"log/slog"
	"net/http"
)

type Handlers struct {
	UserHandler *UserHandler
}

type Response struct {
	Error string `json:"error,omitempty"`
	Data  any    `json:"data,omitempty"`
}

func sendJSON(w http.ResponseWriter, resp Response, status int) {
	w.Header().Set("Content-Type", "application/json")

	data, err := json.Marshal(resp)
	if err != nil {
		slog.Error("failed to marshal response", "error", err)
		sendJSON(w, Response{Error: "something went wrong"}, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(status)
	if _, err := w.Write(data); err != nil {
		slog.Error("failed to write response to client", "error", err)
		return
	}
}

func NewHandlers(repo *repositories.Repo) *Handlers {
	return &Handlers{
		UserHandler: NewUserHandler(repo.UserRepo),
	}
}
