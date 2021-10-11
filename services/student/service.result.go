package student

import (
	"github.com/restuwahyu13/gin-rest-api/models"
	repositorys "github.com/restuwahyu13/gin-rest-api/repositorys/student"
	"github.com/restuwahyu13/gin-rest-api/schemas"
)

type ServiceResult interface {
	ResultStudentService(input *schemas.SchemaStudent) (*models.ModelStudent, schemas.SchemaDatabaseError)
}

type serviceResult struct {
	repository repositorys.RepositoryResult
}

func NewServiceResult(repository repositorys.RepositoryResult) *serviceResult {
	return &serviceResult{repository: repository}
}

func (s *serviceResult) ResultStudentService(input *schemas.SchemaStudent) (*models.ModelStudent, schemas.SchemaDatabaseError) {

	var student schemas.SchemaStudent
	student.ID = input.ID

	res, err := s.repository.ResultStudentRepository(&student)
	return res, err
}
