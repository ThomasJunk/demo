package server

//AddRoutes to server
func (s *Server) AddRoutes() {
	s.Router.Get("/", s.Home)
	s.Router.Get("/login", s.Login)
	s.Router.Get("/logout", s.Authenticated(s.Logout))
	s.Router.Get("/content", s.Authenticated(s.Content))
}
