package config

import (
	"log"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Auth0 Auth0Config `json:"auth0" mapstructure:"auth0"`
}

type Auth0Config struct {
	Domain       string `json:"domain" mapstructure:"domain"`
	ClientId     string `json:"client_id" mapstructure:"client_id"`
	ClientSecret string `json:"client_secret" mapstructure:"client_secret"`
	CallbackUrl  string `json:"callback_url" mapstructure:"callback_url"`
}

func LoadConfig() *Config {
	cfg := &Config{}

	v := viper.NewWithOptions(viper.KeyDelimiter("__"))
	v.SetConfigFile(".env")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "__"))
	v.AutomaticEnv()
	v.AddConfigPath(".")

	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("Failed to read config %v", err)
	}

	err := v.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}

	return cfg
}
