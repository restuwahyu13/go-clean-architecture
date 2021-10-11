package services

import (
	"github.com/restuwahyu13/gin-rest-api/models"
	repositorys "github.com/restuwahyu13/gin-rest-api/repositorys/auth"
	"github.com/restuwahyu13/gin-rest-api/schemas"
)

type ServiceForgot interface {
	ForgotService(input *schemas.SchemaAuth) (*models.ModelAuth, schemas.SchemaDatabaseError)
}

type serviceForgot struct {
	repository repositorys.RepositoryForgot
}

func NewServiceForgot(repository repositorys.RepositoryForgot) *serviceForgot {
	return &serviceForgot{repository: repository}
}

func (s *serviceForgot) ForgotService(input *schemas.SchemaAuth) (*models.ModelAuth, schemas.SchemaDatabaseError) {

	var schema schemas.SchemaAuth
	schema.Email = input.Email

	res, err := s.repository.ForgotRepository(&schema)
	return res, err
}
