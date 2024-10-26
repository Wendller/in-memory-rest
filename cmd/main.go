package main

import (
	"github.com/in-memory-rest/configs"
	"log/slog"
	"net/http"
)

func main() {
	cfg := configs.LoadConfig()

	if err := http.ListenAndServe("localhost:8080", cfg.Router); err != nil {
		slog.Error("application initialize error", "error", err)
		return
	}

	slog.Info("all systems initialized")
}
