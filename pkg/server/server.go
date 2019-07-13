package server

import (
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi"
)

//Server structure
type Server struct {
	Router  *chi.Mux
	Session *scs.SessionManager
}

//New instantiates new server
func New(session *scs.SessionManager) {
	s := prepareServer(session)
	srv := &http.Server{
		Handler: s.Session.LoadAndSave(s.Router),
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}

func prepareServer(session *scs.SessionManager) *Server {
	r := chi.NewRouter()
	sessionManager := session
	s := &Server{
		Router:  r,
		Session: sessionManager,
	}
	s.AddMiddleware()
	s.AddRoutes()
	return s
}
