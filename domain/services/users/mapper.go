package users_service

import (
	entitie "github.com/restuwahyu13/go-clean-architecture/domain/entities"
	cons "github.com/restuwahyu13/go-clean-architecture/shared/constants"
	"github.com/restuwahyu13/go-clean-architecture/shared/dto"
)

func toLoginMapper(enttie *entitie.UsersEntitie) dto.Users {
	return dto.Users{
		ID:        enttie.ID,
		Name:      enttie.Name,
		Status:    enttie.Status,
		CreatedAt: enttie.CreatedAt.Format(cons.DATE_TIME_FORMAT),
	}
}
