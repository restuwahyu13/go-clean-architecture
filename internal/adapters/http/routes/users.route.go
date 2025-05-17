package route

import (
	"github.com/go-chi/chi/v5"
	"github.com/restuwahyu13/go-clean-architecture/shared/dto"
	helper "github.com/restuwahyu13/go-clean-architecture/shared/helpers"
	inf "github.com/restuwahyu13/go-clean-architecture/shared/interfaces"
)

type usersRoute struct {
	router     chi.Router
	controller inf.IUsersController
}

func NewUsersRoute(options dto.RouteOptions[inf.IUsersController]) {
	route := usersRoute{router: options.ROUTER, controller: options.CONTROLLER}

	route.router.Route(helper.Version("users"), func(r chi.Router) {
		r.Get("/", route.controller.Ping)
	})
}
