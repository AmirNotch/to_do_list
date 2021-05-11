package to_do_list

import (
	"context"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(post string) error {
	s.httpServer = &http.Server{
		Addr:           ":" + post,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    1 * time.Second,
		WriteTimeout:   1 * time.Second,
	}

	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}