package dto

import (
	"github.com/go-chi/chi/v5"
)

type GracefulConfig struct {
	HANDLER *chi.Mux
	ENV     Environtment
}
