package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("test"))

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
	w.Header().Set("Content-Type", "application/json")
	var order OrderCheck
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	session, err := store.Get(r, "session")
	if err != nil {
		log.Print(err)
	}
	session.Values["OrderUID"] = order.UID
	err = session.Save(r, w)
	if err != nil {
		log.Print(err)
	}
	fmt.Fprint(w, "{}")
}

func (h *Handler) result(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "session")
	if err != nil {
		log.Print(err)
	}
	order := session.Values["OrderUID"]
	uid := order.(string)
	val, err := Cache.Get(uid)
	if err != nil {
		log.Println(err)
	}
	ord := OrderAnother{}
	err = json.Unmarshal(val, &ord)
	if err != nil {
		log.Print(err)
	}
	tmp, err := template.ParseFiles("result.html")
	if err != nil {
		fmt.Println(err)
	}

	err = tmp.Execute(w, ord)
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
