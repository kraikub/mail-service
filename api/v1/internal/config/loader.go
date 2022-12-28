package config

import (
	"os"

	"github.com/spf13/viper"
)

var EnvList = []string{
	"SERVER_NAME",
	"SERVER_PORT",
	"SMTP_HOST",
	"SMTP_PORT",
	"SMTP_SERVICE_EMAIL",
	"SMTP_SERVICE_EMAIL_PASSWORD",
}

type Config struct {
	Server server `mapstructure:",squash"`
	Smtp   smtp   `mapstructure:",squash"`
}

type server struct {
	Name string `mapstructure:"SERVER_NAME"`
	Port int    `mapstructure:"SERVER_PORT"`
}

type smtp struct {
	Host                 string `mapstructure:"SMTP_HOST"`
	Port                 int    `mapstructure:"SMTP_PORT"`
	ServiceEmail         string `mapstructure:"SMTP_SERVICE_EMAIL"`
	ServiceEmailPassword string `mapstructure:"SMTP_SERVICE_EMAIL_PASSWORD"`
}

func getProductionRuntimeConfig() (Config, error) {
	var conf Config
	viper.AutomaticEnv()
	err := viper.Unmarshal(&conf)
	return conf, err
}

func GetRuntimeConfig() (Config, error) {

	if os.Getenv("KRAIKUB_ENV") == "production" {
		return getProductionRuntimeConfig()
	}
	var conf Config
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		return Config{}, err
	}

	err = viper.Unmarshal(&conf)
	return conf, err
}
