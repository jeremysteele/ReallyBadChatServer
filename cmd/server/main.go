package main

import (
	"github.com/jeremysteele/reallybadchatserver/pkg/config"
	"github.com/jeremysteele/reallybadchatserver/pkg/server"
	log "github.com/sirupsen/logrus"
	"os"
)

func initLogger() {
	log.SetLevel(log.DebugLevel)
	log.SetOutput(os.Stdout)
	log.SetFormatter(&log.JSONFormatter{})
}

func main() {
	initLogger()

	c := config.ReadConfig()

	s := server.NewServer(c)

	log.Infof("Listening on port %d", c.ServerPort)

	if err := s.Run(); err != nil {
		log.WithError(err).Fatal("critical error while running server")
	}
}