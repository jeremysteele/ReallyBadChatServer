package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	ServerPort int
}

func ReadConfig() *Config {
	c := new(Config)

	viper.SetDefault("ServerPort", 1234)
	viper.AutomaticEnv()
	viper.SetEnvPrefix("RBCS")

	err := viper.Unmarshal(c)
	if err != nil {
		log.WithError(err).Fatal("unable to decode into struct")
	}

	return c
}