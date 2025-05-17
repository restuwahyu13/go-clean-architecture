package dto

import (
	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

type (
	ServiceOptions struct {
		ENV Environtment
		DB  *sqlx.DB
		RDS *redis.Client
	}

	UsecaseOptions[T any] struct {
		SERVICE T
	}

	ControllerOptions[T any] struct {
		USECASE T
	}

	RouteOptions[T any] struct {
		ENV        Environtment
		RDS        *redis.Client
		ROUTER     chi.Router
		CONTROLLER T
	}

	ModuleOptions struct {
		ENV    Environtment
		DB     *sqlx.DB
		RDS    *redis.Client
		ROUTER chi.Router
	}
)
