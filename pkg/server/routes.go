package server

import "github.com/ThomasJunk/demo/pkg/controller"

//AddRoutes to server
func (s *Server) AddRoutes() {
	c := controller.New(s.Configuration)
	s.Router.Get("/", s.Home)
	s.Router.Get("/api/login", c.Login)
	s.Router.Get("/api/logout", s.Authenticated(c.Logout))
	s.Router.Get("/api/content", s.Authenticated(c.Content))
}
