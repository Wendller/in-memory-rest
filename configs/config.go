package configs

import (
	"github.com/go-chi/chi/v5"
	"github.com/in-memory-rest/internal/database"
	"github.com/in-memory-rest/internal/database/repositories"
	"github.com/in-memory-rest/internal/web"
	"github.com/in-memory-rest/internal/web/handlers"
)

type Config struct {
	DB      *database.DB
	Router  *chi.Mux
	Handler *handlers.Handlers
}

func LoadConfig() *Config {
	db := database.NewDB()
	repo := repositories.NewRepo(db)
	h := handlers.NewHandlers(repo)

	r := chi.NewRouter()
	web.SetupRoutes(r, h)

	return &Config{
		DB:      db,
		Router:  r,
		Handler: h,
	}
}
