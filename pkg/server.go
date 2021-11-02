package pkg

import (
	"context"
	"net/http"
)

type Config struct {
	Port string `yaml:"port" env:"PORT"`
	Host string `yaml:"host" env:"HOST" env-default:"0.0.0.0"`
}
type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(path string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:    path,
		Handler: handler,
	}
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
