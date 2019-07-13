package server

import (
	"net/http"
)

//Content serves content
func (s *Server) Content(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Protected"))
}
