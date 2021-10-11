package services

import (
	"github.com/restuwahyu13/gin-rest-api/models"
	repositorys "github.com/restuwahyu13/gin-rest-api/repositorys/auth"
	"github.com/restuwahyu13/gin-rest-api/schemas"
)

type ServiceResend interface {
	ResendService(input *schemas.SchemaAuth) (*models.ModelAuth, schemas.SchemaDatabaseError)
}

type serviceResend struct {
	repository repositorys.RepositoryResend
}

func NewServiceResend(repository repositorys.RepositoryResend) *serviceResend {
	return &serviceResend{repository: repository}
}

func (s *serviceResend) ResendService(input *schemas.SchemaAuth) (*models.ModelAuth, schemas.SchemaDatabaseError) {

	var schema schemas.SchemaAuth
	schema.Email = input.Email

	res, err := s.repository.ResendRepository(&schema)
	return res, err
}
