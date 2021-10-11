package student

import (
	"github.com/restuwahyu13/gin-rest-api/models"
	repositorys "github.com/restuwahyu13/gin-rest-api/repositorys/student"
	"github.com/restuwahyu13/gin-rest-api/schemas"
)

type ServiceDelete interface {
	DeleteStudentService(input *schemas.SchemaStudent) (*models.ModelStudent, schemas.SchemaDatabaseError)
}

type serviceDelete struct {
	repository repositorys.RepositoryDelete
}

func NewServiceDelete(repository repositorys.RepositoryDelete) *serviceDelete {
	return &serviceDelete{repository: repository}
}

func (s *serviceDelete) DeleteStudentService(input *schemas.SchemaStudent) (*models.ModelStudent, schemas.SchemaDatabaseError) {

	var student schemas.SchemaStudent
	student.ID = input.ID

	res, err := s.repository.DeleteStudentRepository(&student)
	return res, err
}
