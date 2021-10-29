package pkg

import (
	"context"
	"net/http"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port string) error {
	s.httpServer = &http.Server{
		Addr: ":" + port,
	}
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
