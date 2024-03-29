package server

import (
	"net/http"
	"time"

	chilogger "github.com/766b/chi-logger"
	authz "github.com/casbin/chi-authz"
	"github.com/go-chi/chi/middleware"
	"github.com/unrolled/secure"
)

//AddMiddleware to server
func (s *Server) AddMiddleware() {
	secureMiddleware := secure.New(secure.Options{
		FrameDeny: true,
	})

	s.Router.Use(middleware.RequestID)
	s.Router.Use(middleware.RealIP)
	s.Router.Use(chilogger.NewZapMiddleware("router", s.Log))
	s.Router.Use(middleware.Recoverer)
	s.Router.Use(secureMiddleware.Handler)
	s.Router.Use(authz.Authorizer(s.Enforcer))
	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	s.Router.Use(middleware.Timeout(60 * time.Second))
}

//Authenticated routes
func (s *Server) Authenticated(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		msg := s.Session.GetString(r.Context(), "userid")
		if msg == "" {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}
		next(w, r)
	}
}

//Authorized middleware
func (s *Server) Authorized(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := s.Session.GetString(r.Context(), "userid")
		method := r.Method
		path := r.URL.Path
		if s.Enforcer.Enforce(user, path, method) {
			next(w, r)
		} else {
			http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
		}
	}
}
