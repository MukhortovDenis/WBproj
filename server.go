package main

import (
	"context"
	"fmt"
	"net/http"
	"text/template"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(path string, router http.Handler) error {
	s.httpServer = &http.Server{
		Addr:    path,
		Handler: router,
	}
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}

func NewRouter() *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	return router
}

func (h *Handler) index(w http.ResponseWriter, r *http.Request) {
	tmp, err := template.ParseFiles("index.html")
	if err != nil {
		fmt.Println(err)
	}
	err = tmp.Execute(w, nil)
	if err != nil {
		fmt.Fprint(w, err)
	}
}

func (h *Handler) checkOrder(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "eee бой")
}

func (h *Handler) result(w http.ResponseWriter, r *http.Request){
	tmp, err := template.ParseFiles("result.html")
	if err != nil {
		fmt.Println(err)
	}
	order := OrderAnother{}
	err = tmp.Execute(w, order)
	if err != nil {
		fmt.Fprint(w, err)
	}
}

func (h *Handler) mainHandle() *chi.Mux {
	router := NewRouter()
	fileServer(router)
	router.Get("/", h.index)
	router.Get("/result", h.result)
	router.Post("/check-order", h.checkOrder)
	return router
}
