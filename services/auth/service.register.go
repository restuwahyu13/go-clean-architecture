package services

import (
	"github.com/restuwahyu13/gin-rest-api/models"
	repositorys "github.com/restuwahyu13/gin-rest-api/repositorys/auth"
	"github.com/restuwahyu13/gin-rest-api/schemas"
)

type ServiceRegister interface {
	RegisterService(input *schemas.SchemaAuth) (*models.ModelAuth, schemas.SchemaDatabaseError)
}

type serviceRegister struct {
	repository repositorys.RepositoryRegister
}

func NewServiceRegister(repository repositorys.RepositoryRegister) *serviceRegister {
	return &serviceRegister{repository: repository}
}

func (s *serviceRegister) RegisterService(input *schemas.SchemaAuth) (*models.ModelAuth, schemas.SchemaDatabaseError) {

	var schema schemas.SchemaAuth
	schema.Fullname = input.Fullname
	schema.Email = input.Email
	schema.Password = input.Password

	res, err := s.repository.RegisterRepository(&schema)
	return res, err
}
