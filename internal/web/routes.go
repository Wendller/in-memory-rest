package web

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/in-memory-rest/internal/web/handlers"
	"net/http"
)

func SetupRoutes(router *chi.Mux, handlers *handlers.Handlers) {
	router.Use(middleware.Logger)
	router.Use(middleware.RequestID)
	router.Use(middleware.Recoverer)

	router.Route("/api", func(r chi.Router) {
		r.Mount("/users", userRouter(*handlers.UserHandler))
	})

}

func userRouter(userHandler handlers.UserHandler) http.Handler {
	r := chi.NewRouter()
	r.Get("/", userHandler.GetAllUsers)
	r.Get("/{id}", userHandler.GetUserById)
	r.Post("/", userHandler.Insert)
	r.Put("/{id}", userHandler.Update)
	r.Delete("/{id}", userHandler.DeleteUserById)

	return r
}
