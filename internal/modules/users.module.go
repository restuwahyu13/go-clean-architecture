package module

import (
	users_service "github.com/restuwahyu13/go-clean-architecture/domain/services/users"
	controller "github.com/restuwahyu13/go-clean-architecture/internal/adapters/http/controllers"
	route "github.com/restuwahyu13/go-clean-architecture/internal/adapters/http/routes"
	"github.com/restuwahyu13/go-clean-architecture/shared/dto"
	inf "github.com/restuwahyu13/go-clean-architecture/shared/interfaces"
	usecase "github.com/restuwahyu13/go-clean-architecture/usecases"
)

type usersModule[IService any] struct {
	service IService
}

func NewUsersModule[IService any](options dto.ModuleOptions) inf.IUsersModule[IService] {
	service := users_service.NewUsersService(dto.ServiceOptions{ENV: options.ENV, DB: options.DB, RDS: options.RDS})

	usecase := usecase.NewUsersUsecase(dto.UsecaseOptions[inf.IUsersService]{SERVICE: service})

	controller := controller.NewUsersController(dto.ControllerOptions[inf.IUsersUsecase]{USECASE: usecase})

	route.NewUsersRoute(dto.RouteOptions[inf.IUsersController]{ROUTER: options.ROUTER, CONTROLLER: controller})

	return usersModule[IService]{service: any(service).(IService)}
}

func (m usersModule[IService]) Service() IService {
	return m.service
}
