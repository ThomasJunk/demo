package main

import (
	"log"
	"time"

	"bitbucket.com/ThomasJunk/demo/pkg/server"
	"github.com/alexedwards/scs/boltstore"
	"github.com/alexedwards/scs/v2"
	"go.etcd.io/bbolt"
)

func main() {
	db, err := bbolt.Open("/tmp/bolt.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	sessionManager := scs.New()
	sessionManager.Store = boltstore.NewWithCleanupInterval(db, 20*time.Second)
	sessionManager.Lifetime = time.Minute
	server.New(sessionManager)

}
