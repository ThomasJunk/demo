package controller

import (
	"github.com/alexedwards/scs/v2"
)

//Handler for controllers
type Handler struct {
	Session *scs.SessionManager
}

//New generates a new Handler
func New(s *scs.SessionManager) *Handler {
	return &Handler{
		Session: s,
	}
}
