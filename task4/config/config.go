package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	JWTSecret string `mapstructure:"JWT_SECRET"`
}

var AppCconfig *Config

func LoadConfig() {
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Error reading config file, %s", err)
	}
	AppCconfig = &Config{
		JWTSecret: viper.GetString("JWT_SECRET"),
	}
}
