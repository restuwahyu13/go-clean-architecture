package services

import (
	"github.com/restuwahyu13/gin-rest-api/models"
	repositorys "github.com/restuwahyu13/gin-rest-api/repositorys/auth"
	"github.com/restuwahyu13/gin-rest-api/schemas"
)

type ServiceLogin interface {
	LoginService(input *schemas.SchemaAuth) (*models.ModelAuth, schemas.SchemaDatabaseError)
}

type serviceLogin struct {
	repository repositorys.RepositoryLogin
}

func NewServiceLogin(repository repositorys.RepositoryLogin) *serviceLogin {
	return &serviceLogin{repository: repository}
}

func (s *serviceLogin) LoginService(input *schemas.SchemaAuth) (*models.ModelAuth, schemas.SchemaDatabaseError) {

	var schema schemas.SchemaAuth
	schema.Email = input.Email
	schema.Password = input.Password

	res, err := s.repository.LoginRepository(&schema)
	return res, err
}
