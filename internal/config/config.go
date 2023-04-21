package config

import (
	"errors"
	"os"
	"strconv"
)

type Config struct {
	Port int    `env:"PORT,required"`
	Dsn  string `env:"DSN,required"`
}

func New() (*Config, error) {
	port := os.Getenv("PORT")
	dsn := os.Getenv("DSN")
	migrationUrl := os.Getenv("MIGRATION_URL")
	if port == "" && dsn == "" && migrationUrl == "" {
		return nil, errors.New("no data in env")
	}

	p, err := strconv.Atoi(port)
	if err != nil {
		return nil, err
	}

	return &Config{
		Port: p,
		Dsn:  dsn,
	}, nil
}
