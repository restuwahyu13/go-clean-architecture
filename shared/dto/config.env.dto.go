package dto

import opt "github.com/restuwahyu13/go-clean-architecture/shared/output"

type Config struct {
	ENV            string `env:"GO_ENV" mapstructure:"GO_ENV"`
	PORT           string `env:"PORT" mapstructure:"PORT"`
	INBOUND_SIZE   int    `env:"INBOUND_SIZE" mapstructure:"INBOUND_SIZE"`
	DSN            string `env:"PG_DSN" mapstructure:"PG_DSN"`
	CSN            string `env:"REDIS_CSN" mapstructure:"REDIS_CSN"`
	JWT_SECRET_KEY string `env:"JWT_SECRET_KEY" mapstructure:"JWT_SECRET_KEY"`
	JWT_EXPIRED    int    `env:"JWT_EXPIRED" mapstructure:"JWT_EXPIRED"`
}

type (
	Environtment struct {
		APP      *opt.Application
		REDIS    *opt.Redis
		POSTGRES *opt.Postgres
		JWT      *opt.Jwt
	}
)
