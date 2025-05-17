package main

import (
	"compress/zlib"
	"context"
	"os"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/jmoiron/sqlx"
	"github.com/oxequa/grace"
	"github.com/redis/go-redis/v9"
	config "github.com/restuwahyu13/go-clean-architecture/configs"
	con "github.com/restuwahyu13/go-clean-architecture/internal/infrastructure/connections"
	module "github.com/restuwahyu13/go-clean-architecture/internal/modules"
	cons "github.com/restuwahyu13/go-clean-architecture/shared/constants"
	"github.com/restuwahyu13/go-clean-architecture/shared/dto"
	helper "github.com/restuwahyu13/go-clean-architecture/shared/helpers"
	inf "github.com/restuwahyu13/go-clean-architecture/shared/interfaces"
	"github.com/restuwahyu13/go-clean-architecture/shared/pkg"
	"github.com/unrolled/secure"
)

type (
	IApi interface {
		Middleware()
		Module()
		Listener()
	}

	Api struct {
		ENV    dto.Environtment
		ROUTER *chi.Mux
		DB     *sqlx.DB
		RDS    *redis.Client
	}
)

var (
	err error
	env dto.Environtment
)

func init() {
	transform := helper.NewTransform()

	env_res, err := config.NewEnvirontment(".env", ".", "env")
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

func main() {
	ctx := context.Background()
	router := chi.NewRouter()

	db, err := con.SqlConnection(ctx, env)
	if err != nil {
		pkg.Logrus("fatal", err)
		return
	}

	rds, err := con.RedisConnection(env)
	if err != nil {
		pkg.Logrus("fatal", err)
		return
	}

	app := NewApi(Api{ENV: env, ROUTER: router, DB: db, RDS: rds})
	app.Middleware()
	app.Module()
	app.Listener()
}

func NewApi(options Api) IApi {
	return Api{
		ENV:    options.ENV,
		ROUTER: options.ROUTER,
		DB:     options.DB,
		RDS:    options.RDS,
	}
}

func (i Api) Middleware() {
	if i.ENV.APP.ENV != cons.PROD {
		i.ROUTER.Use(middleware.Logger)
	}

	i.ROUTER.Use(middleware.Recoverer)
	i.ROUTER.Use(middleware.RealIP)
	i.ROUTER.Use(middleware.NoCache)
	i.ROUTER.Use(middleware.GetHead)
	i.ROUTER.Use(middleware.Compress(zlib.BestCompression))
	i.ROUTER.Use(middleware.AllowContentType("application/json"))
	i.ROUTER.Use(cors.Handler(cors.Options{
		AllowedOrigins:     []string{"*"},
		AllowedMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowedHeaders:     []string{"Accept", "Content-Type", "Authorization"},
		AllowCredentials:   true,
		OptionsPassthrough: true,
		MaxAge:             900,
	}))
	i.ROUTER.Use(secure.New(secure.Options{
		FrameDeny:            true,
		ContentTypeNosniff:   true,
		BrowserXssFilter:     true,
		STSIncludeSubdomains: true,
		STSPreload:           true,
		STSSeconds:           900,
	}).Handler)
}

func (i Api) Module() {
	module.NewUsersModule[inf.IUsersService](dto.ModuleOptions{
		ENV:    i.ENV,
		DB:     i.DB,
		RDS:    i.RDS,
		ROUTER: i.ROUTER,
	})
}

func (i Api) Listener() {
	err := pkg.Graceful(func() *dto.GracefulConfig {
		return &dto.GracefulConfig{HANDLER: i.ROUTER, ENV: i.ENV}
	})

	recover := grace.Recover(&err)
	recover.Stack()

	if err != nil {
		pkg.Logrus("fatal", err)
		os.Exit(1)
	}
}
