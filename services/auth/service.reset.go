package services

import (
	"github.com/restuwahyu13/gin-rest-api/models"
	repositorys "github.com/restuwahyu13/gin-rest-api/repositorys/auth"
	"github.com/restuwahyu13/gin-rest-api/schemas"
)

type ServiceReset interface {
	ResetService(input *schemas.SchemaAuth) (*models.ModelAuth, schemas.SchemaDatabaseError)
}

type serviceReset struct {
	repository repositorys.RepositoryReset
}

func NewServiceReset(repository repositorys.RepositoryReset) *serviceReset {
	return &serviceReset{repository: repository}
}

func (s *serviceReset) ResetService(input *schemas.SchemaAuth) (*models.ModelAuth, schemas.SchemaDatabaseError) {

	var schema schemas.SchemaAuth
	schema.Email = input.Email
	schema.Password = input.Email
	schema.Cpassword = input.Cpassword
	schema.Active = input.Active

	res, err := s.repository.ResetRepository(&schema)
	return res, err
}
