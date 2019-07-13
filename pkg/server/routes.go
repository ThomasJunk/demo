package server

import "bitbucket.com/ThomasJunk/demo/pkg/controller"

//AddRoutes to server
func (s *Server) AddRoutes() {
	h := controller.New(s.Session)
	s.Router.Get("/", s.Home)
	s.Router.Get("/login", h.Login)
	s.Router.Get("/logout", s.Authenticated(h.Logout))
	s.Router.Get("/content", s.Authenticated(h.Content))
}
