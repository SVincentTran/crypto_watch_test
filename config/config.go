package config

import (
	"log"
	"os"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Auth0    Auth0Config      `json:"auth0" mapstructure:"auth0"`
	Postgres PostgresqlConfig `json:"postgres" mapstructure:"postgres"`
}

type Auth0Config struct {
	Domain       string `json:"domain" mapstructure:"domain"`
	ClientId     string `json:"client_id" mapstructure:"client_id"`
	ClientSecret string `json:"client_secret" mapstructure:"client_secret"`
	CallbackUrl  string `json:"callback_url" mapstructure:"callback_url"`
}

type PostgresqlConfig struct {
	User     string `json:"user" mapstructure:"user"`
	Password string `json:"password" mapstructure:"password"`
	DB       string `json:"db" mapstructure:"db"`
	Host     string `json:"host" mapstructure:"host"`
	Port     string `json:"port" mapstructure:"port"`
}

func LoadConfig() *Config {
	cfg := &Config{}

	v := viper.NewWithOptions(viper.KeyDelimiter("__"))
	v.SetConfigFile(".env")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "__"))
	v.AutomaticEnv()
	v.AddConfigPath(".")

	err := v.ReadInConfig()
	if err == nil {
		err := v.Unmarshal(&cfg)
		if err != nil {
			panic(err)
		}

		return cfg
	}

	log.Printf("Failed to read config %v", err)

	domain := os.Getenv("AUTH0__DOMAIN")
	clientId := os.Getenv("AUTH0__CLIENT_ID")
	clientSecret := os.Getenv("AUTH0__CLIENT_SECRET")
	callbackUrl := os.Getenv("AUTH0__CALLBACK_URL")

	dbUser := os.Getenv("POSTGRES__USER")
	dbPassword := os.Getenv("POSTGRES__PASSWORD")
	dbName := os.Getenv("POSTGRES__DB")
	dbHost := os.Getenv("POSTGRES__HOST")
	dbPort := os.Getenv("POSTGRES__PORT")

	if domain == "" || clientId == "" || clientSecret == "" || callbackUrl == "" {
		log.Fatal("Missing one or more environment variables!!!")
	}

	cfg.Auth0.Domain = domain
	cfg.Auth0.ClientId = clientId
	cfg.Auth0.ClientSecret = clientSecret
	cfg.Auth0.CallbackUrl = callbackUrl

	cfg.Postgres.User = dbUser
	cfg.Postgres.Password = dbPassword
	cfg.Postgres.DB = dbName
	cfg.Postgres.Host = dbHost
	cfg.Postgres.Port = dbPort

	return cfg
}
