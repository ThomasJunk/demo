package controller

import (
	"net/http"
)

//Content serves content
func (h *Handler) Content(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Protected"))
}
