package pkg

import (
	"WBproj/cmd/service"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler{
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Get("/", h.signIn)
	router.Get("/sign-in", h.signIn)
	router.Post("/sign-up", h.signUp)
	return router
}
