package controller

import "net/http"

//Login handles login
func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	h.Session.Put(r.Context(), "userid", "1234")
	w.Write([]byte("Logged in"))
}

//Logout handles logout
func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	h.Session.Destroy(r.Context())
	w.Write([]byte("Logged out"))
}
