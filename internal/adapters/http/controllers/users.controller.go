package controller

import (
	"net/http"

	cons "github.com/restuwahyu13/go-clean-architecture/shared/constants"
	"github.com/restuwahyu13/go-clean-architecture/shared/dto"
	helper "github.com/restuwahyu13/go-clean-architecture/shared/helpers"
	inf "github.com/restuwahyu13/go-clean-architecture/shared/interfaces"
	opt "github.com/restuwahyu13/go-clean-architecture/shared/output"
	"github.com/restuwahyu13/go-clean-architecture/shared/pkg"
)

type usersController struct {
	usecase inf.IUsersUsecase
}

func NewUsersController(options dto.ControllerOptions[inf.IUsersUsecase]) inf.IUsersController {
	return &usersController{usecase: options.USECASE}
}

func (c usersController) Ping(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	res := opt.Response{}

	if res = c.usecase.Ping(ctx); res.StatCode >= http.StatusBadRequest {
		if res.StatCode >= http.StatusInternalServerError {
			pkg.Logrus(cons.ERROR, res.ErrMsg)
			res.ErrMsg = cons.DEFAULT_ERR_MSG
		}

		helper.Api(rw, r, res)
		return
	}

	helper.Api(rw, r, res)
	return
}
