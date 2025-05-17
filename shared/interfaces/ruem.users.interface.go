package inf

import (
	"context"
	"net/http"

	entitie "github.com/restuwahyu13/go-clean-architecture/domain/entities"
	opt "github.com/restuwahyu13/go-clean-architecture/shared/output"
)

type (
	IUsersRepositorie interface {
		Find(query string, args ...any) ([]entitie.UsersEntitie, error)
		FindOne(query string, args ...any) (*entitie.UsersEntitie, error)
		Create(dest any, query string, args ...any) error
		Update(dest any, query string, args ...any) error
		Delete(dest any, query string, args ...any) error
	}

	IUsersService interface {
		Ping(ctx context.Context) opt.Response
	}

	IUsersException interface {
		Login(key string) string
	}

	IUsersUsecase interface {
		Ping(ctx context.Context) opt.Response
	}

	IUsersController interface {
		Ping(rw http.ResponseWriter, r *http.Request)
	}

	IUsersModule[IUserService any] interface {
		Service() IUserService
	}
)
