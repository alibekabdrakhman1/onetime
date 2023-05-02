package config

import (
	"errors"
	"github.com/caarlos0/env/v6"
)

type Config struct {
	HTTP     HTTP
	Database Database
	JwtKey   string `envDefault:"your-256-bit-secret"`
}
type HTTP struct {
	PORT string `env:"PORT" envDefault:"8586"`
	URL  string `env:"URL" envDefault:"localhost"`
}
type Database struct {
	URL string `env:"URL" envDefault:"mongodb+srv://alibekabdrakhman:R6J9M97WIKwcnIqe@ontime.lb54dsn.mongodb.net/test"`
}

func New() (*Config, error) {
	cfg := Config{JwtKey: "your-256-bit-secret"}
	if err := env.Parse(&cfg); err != nil {
		return nil, errors.New("cfg not created")
	}
	return &cfg, nil
}
