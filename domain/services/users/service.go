package users_service

import (
	"context"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
	"github.com/restuwahyu13/go-clean-architecture/shared/dto"
	inf "github.com/restuwahyu13/go-clean-architecture/shared/interfaces"
	opt "github.com/restuwahyu13/go-clean-architecture/shared/output"
)

type usersService struct {
	env dto.Environtment
	db  *sqlx.DB
	rds *redis.Client
}

func NewUsersService(options dto.ServiceOptions) inf.IUsersService {
	return &usersService{env: options.ENV, db: options.DB, rds: options.RDS}
}

func (s usersService) Ping(ctx context.Context) opt.Response {
	res := opt.Response{}

	res.StatCode = http.StatusOK
	res.Message = "Ping!"

	return res
}
