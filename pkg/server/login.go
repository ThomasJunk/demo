package server

import "net/http"

//Login handles login
func (s *Server) Login(w http.ResponseWriter, r *http.Request) {
	s.Session.Put(r.Context(), "userid", "1234")
	w.Write([]byte("Logged in"))
}

//Logout handles logout
func (s *Server) Logout(w http.ResponseWriter, r *http.Request) {
	s.Session.Destroy(r.Context())
	w.Write([]byte("Logged out"))
}
