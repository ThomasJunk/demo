package server

import "net/http"

//Home route
func (s *Server) Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome"))
}
