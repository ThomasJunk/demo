package server

import (
	"net/http"
)

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

//AddRoutes to server
func (s *Server) AddRoutes() {
	s.Router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	s.Router.Get("/login", s.Login)
	s.Router.Get("/logout", s.Logout)
	s.Router.Get("/content", s.Authenticated(s.Content))
}
