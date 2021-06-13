package config

import (
	"fmt"
	"os"

	"github.com/kelseyhightower/envconfig"

	"github.com/joho/godotenv"
)

type Config struct {
	Environment                   string `envconfig:"APP_ENVIRONMENT" default:"development"`
	Host                          string `envconfig:"HOST" default:"localhost"`
	HTTPPort                      int    `envconfig:"HTTP_PORT" required:"true" default:"8080"`
	GRPCPort                      int    `envconfig:"GRPC_PORT" required:"true" default:"9090"`
	HttpClientTimeout             int    `envconfig:"HTTP_CLIENT_TIMEOUT"`
	LogLevel                      int    `envconfig:"LOG_LEVEL" default:"-1"`
	LogTimeFormat                 string `envconfig:"LOG_TIME_FORMAT" default:"2006-01-02T15:04:05.999999999Z07:00"`
	JwtIssuer                     string `envconfig:"JWT_ISSUER"`
	JwtLeeway                     string `envconfig:"JWT_LEEWAY" default:"60"`
	PostgresUser                  string `envconfig:"POSTGRES_USER"`
	PostgresDatabase              string `envconfig:"POSTGRES_DB"`
	PostgresPassword              string `envconfig:"POSTGRES_PASSWORD"`
	PostgresHost                  string `envconfig:"POSTGRES_HOST" default:"localhost"`
	PostgresPort                  int    `envconfig:"POSTGRES_PORT" default:"5432"`
	PostgresMaxConnections        int    `envconfig:"POSTGRES_MAX_CONNECTIONS" default:"10"`
	PostgresMaxConnectionLifetime int    `envconfig:"POSTGRES_MAX_CONNECTION_LIFETIME" default:"30"`
}

func GetConfig() (*Config, error) {
	if _, err := os.Stat(".env"); err == nil {
		err = godotenv.Load()
		if err != nil {
			fmt.Println("Error loading .env file")
		}
	}

	var config Config
	err := envconfig.Process("", &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func (c *Config) GRPCServerAddress() string {
	return fmt.Sprintf("%v:%v", c.Host, c.GRPCPort)
}

func (c *Config) HTTPServerAddress() string {
	return fmt.Sprintf("%v:%v", c.Host, c.HTTPPort)
}
