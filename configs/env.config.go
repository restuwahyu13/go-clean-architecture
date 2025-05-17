package config

import (
	"os"

	genv "github.com/caarlos0/env"
	"github.com/restuwahyu13/go-clean-architecture/shared/dto"
	opt "github.com/restuwahyu13/go-clean-architecture/shared/output"

	"github.com/spf13/viper"
)

func NewEnvirontment(name, path, ext string) (*opt.Environtment, error) {
	cfg := dto.Config{}

	if _, ok := os.LookupEnv("GO_ENV"); !ok {
		viper.SetConfigName(name)
		viper.SetConfigType(ext)
		viper.AddConfigPath(path)
		viper.AutomaticEnv()

		if err := viper.ReadInConfig(); err != nil {
			return nil, err
		}

		if err := viper.Unmarshal(&cfg); err != nil {
			return nil, err
		}
	} else {
		if err := genv.Parse(&cfg); err != nil {
			return nil, err
		}
	}

	return &opt.Environtment{
		APP: &opt.Application{
			ENV:          cfg.ENV,
			PORT:         cfg.PORT,
			INBOUND_SIZE: cfg.INBOUND_SIZE,
		},
		REDIS: &opt.Redis{
			URL: cfg.CSN,
		},
		POSTGRES: &opt.Postgres{
			URL: cfg.DSN,
		},
		JWT: &opt.Jwt{
			SECRET:  cfg.JWT_SECRET_KEY,
			EXPIRED: cfg.JWT_EXPIRED,
		},
	}, nil
}
