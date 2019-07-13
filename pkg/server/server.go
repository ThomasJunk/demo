package server

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/ThomasJunk/demo/pkg/configuration"
	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi"
	"go.uber.org/zap"
)

//Server structure
type Server struct {
	Router        *chi.Mux
	Session       *scs.SessionManager
	Configuration *configuration.Configuration
	Log           *zap.Logger
}

//Run instantiates new server and runs it
func Run(c *configuration.Configuration) {
	s := prepareServer(c)
	srv := &http.Server{
		Handler:      s.Session.LoadAndSave(s.Router),
		Addr:         c.Listen(),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	c.Logger.Info(fmt.Sprintf("Starting server on %s", c.Listen()))
	log.Fatal(srv.ListenAndServe())
}

func prepareServer(c *configuration.Configuration) *Server {
	r := chi.NewRouter()
	sessionManager := c.Session
	s := &Server{
		Router:        r,
		Session:       sessionManager,
		Configuration: c,
		Log:           c.Logger,
	}
	s.AddMiddleware()
	s.AddRoutes()
	return s
}
