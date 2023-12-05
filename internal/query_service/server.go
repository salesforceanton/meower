package query_service

import (
	"context"
	"net/http"
	"time"
)

type QueryServiceServer struct {
	httpServer *http.Server
}

func (s *QueryServiceServer) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	return s.httpServer.ListenAndServe()
}

func (s *QueryServiceServer) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
