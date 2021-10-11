package services

import (
	"github.com/restuwahyu13/gin-rest-api/models"
	repositorys "github.com/restuwahyu13/gin-rest-api/repositorys/auth"
	"github.com/restuwahyu13/gin-rest-api/schemas"
)

type ServiceActivation interface {
	ActivationService(input *schemas.SchemaAuth) (*models.ModelAuth, schemas.SchemaDatabaseError)
}

type serviceActivation struct {
	repository repositorys.RepositoryActivation
}

func NewServiceActivation(repository repositorys.RepositoryActivation) *serviceActivation {
	return &serviceActivation{repository: repository}
}

func (s *serviceActivation) ActivationService(input *schemas.SchemaAuth) (*models.ModelAuth, schemas.SchemaDatabaseError) {

	var schema schemas.SchemaAuth
	schema.Email = input.Email
	schema.Active = input.Active
	schema.Token = input.Token

	res, err := s.repository.ActivationRepository(&schema)
	return res, err
}
