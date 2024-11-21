package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	DB  DB
	API API
}

type DB struct {
	Host               string `envconfig:"DB_HOST" default:"localhost"`
	Port               string `envconfig:"DB_PORT" default:"5432"`
	User               string `envconfig:"DB_USER" default:"postgres"`
	Password           string `envconfig:"DB_PASSWORD" default:"12345"`
	Name               string `envconfig:"DB_NAME" default:"4med"`
	Schema             string `envconfig:"DB_SCHEMA"`
	SSLMode            string `envconfig:"DB_SSL_MODE" default:"disable"`
	MaxOpenConnections *int   `envconfig:"DB_MAX_OPEN_CONNS"`
	MaxIdleConnections *int   `envconfig:"DB_MAX_IDLE_CONNS"`
}

type API struct {
	Address string `envconfig:"LISTEN_ADDR" default:":3001"`
}

func Read() (*Config, error) {
	var c Config
	if err := envconfig.Process("", &c); err != nil {
		return nil, err
	}

	return &c, nil
}
