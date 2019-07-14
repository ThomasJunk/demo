package main

import (
	"log"
	"time"

	"github.com/ThomasJunk/demo/pkg/configuration"
	"github.com/ThomasJunk/demo/pkg/server"
	"github.com/alexedwards/scs/boltstore"
	"github.com/alexedwards/scs/v2"
	"github.com/casbin/casbin"
	"go.etcd.io/bbolt"
	"go.uber.org/zap"
)

func main() {
	db, err := bbolt.Open("/tmp/bolt.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	e := casbin.NewEnforcer("authz_model.conf", "authz_policy.csv")
	s := scs.New()
	s.Store = boltstore.NewWithCleanupInterval(db, 20*time.Second)
	s.Lifetime = time.Minute
	l, _ := zap.NewProduction()
	defer l.Sync()
	c := configuration.New(s, l, e)
	server.Run(c)

}
