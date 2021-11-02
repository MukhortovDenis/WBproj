package pkg

import (
	"WBproj/cmd/service"
	"WBproj/pkg/logging"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
	services *service.Service
	logger logging.Logger
}

func NewHandler(services *service.Service, logger logging.Logger) *Handler {
	return &Handler{
		services: services,
		logger: logger,
	}
}

func (h *Handler) InitRoutes() *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Get("/", h.signIn)
	router.Get("/sign-in", h.signIn)
	router.Get("/sign-up", h.signUp)
	return router
}
