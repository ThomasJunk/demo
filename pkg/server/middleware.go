package server

import (
	"time"

	"github.com/go-chi/chi/middleware"
)

//AddMiddleware to server
func (s *Server) AddMiddleware() {
	// A good base middleware stack
	s.Router.Use(middleware.RequestID)
	s.Router.Use(middleware.RealIP)
	s.Router.Use(middleware.Logger)
	s.Router.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	s.Router.Use(middleware.Timeout(60 * time.Second))
}
