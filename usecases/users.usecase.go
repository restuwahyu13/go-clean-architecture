package usecase

import (
	"context"

	"github.com/restuwahyu13/go-clean-architecture/shared/dto"
	inf "github.com/restuwahyu13/go-clean-architecture/shared/interfaces"
	opt "github.com/restuwahyu13/go-clean-architecture/shared/output"
)

type usersUsecase struct {
	service inf.IUsersService
}

func NewUsersUsecase(options dto.UsecaseOptions[inf.IUsersService]) inf.IUsersUsecase {
	return &usersUsecase{service: options.SERVICE}
}

func (u usersUsecase) Ping(ctx context.Context) opt.Response {
	return u.service.Ping(ctx)
}
