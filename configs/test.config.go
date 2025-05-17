package config

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
	con "github.com/restuwahyu13/go-clean-architecture/internal/infrastructure/connections"
	"github.com/restuwahyu13/go-clean-architecture/shared/dto"
	helper "github.com/restuwahyu13/go-clean-architecture/shared/helpers"
	"github.com/restuwahyu13/go-clean-architecture/shared/pkg"
)

type (
	Test struct {
		CTX context.Context
		ENV dto.Environtment
		DB  *sqlx.DB
		RDS *redis.Client
	}
)

var (
	err error
	env dto.Environtment
)

func init() {
	transform := helper.NewTransform()

	env_res, err := NewEnvirontment(".env", ".", "env")
	if err != nil {
		pkg.Logrus("fatal", err)
		return
	}

	if env_res != nil {
		if err := transform.ResToReq(env_res, &env); err != nil {
			pkg.Logrus("fatal", err)
		}
	}
}

func NewTest() Test {
	ctx := context.Background()

	db, err := con.SqlConnection(ctx, env)
	if err != nil {
		pkg.Logrus("fatal", err)
	}

	rds, err := con.RedisConnection(env)
	if err != nil {
		pkg.Logrus("fatal", err)
	}

	return Test{
		CTX: ctx,
		ENV: env,
		DB:  db,
		RDS: rds,
	}
}
