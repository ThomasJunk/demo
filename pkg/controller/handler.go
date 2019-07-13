package controller

import (
	"github.com/ThomasJunk/demo/pkg/configuration"
	"github.com/alexedwards/scs/v2"
	"go.uber.org/zap"
)

//Handler for controllers
type Handler struct {
	Session       *scs.SessionManager
	Configuration *configuration.Environment
	Log           *zap.Logger
}

//New generates a new Handler
func New(c *configuration.Configuration) *Handler {
	return &Handler{
		Session:       c.Session,
		Configuration: c.Environment,
		Log:           c.Logger,
	}
}
