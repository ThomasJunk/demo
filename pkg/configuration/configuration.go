package configuration

import (
	"fmt"
	"os"

	"github.com/alexedwards/scs/v2"
	"github.com/casbin/casbin"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

//Environment settings
type Environment struct {
	Port    string
	Address string
}

//Configuration struct
type Configuration struct {
	Environment *Environment
	Session     *scs.SessionManager
	Logger      *zap.Logger
	Enforcer    *casbin.Enforcer
}

//New configuration construction
func New(session *scs.SessionManager, l *zap.Logger, e *casbin.Enforcer) *Configuration {
	err := godotenv.Load()
	port := "8000"
	address := "127.0.0.1"
	if err == nil {
		port = os.Getenv("PORT")
		address = os.Getenv("ADDRESS")
	}
	environment := &Environment{
		Port:    port,
		Address: address,
	}
	return &Configuration{
		Environment: environment,
		Session:     session,
		Logger:      l,
		Enforcer:    e,
	}
}

//Listen to address + port
func (c *Configuration) Listen() string {
	return fmt.Sprintf("%s:%s", c.Environment.Address, c.Environment.Port)
}
