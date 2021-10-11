package student

import (
	"github.com/restuwahyu13/gin-rest-api/models"
	model "github.com/restuwahyu13/gin-rest-api/models"
	repositorys "github.com/restuwahyu13/gin-rest-api/repositorys/student"
	"github.com/restuwahyu13/gin-rest-api/schemas"
)

type ServiceResults interface {
	ResultsStudentService() (*[]models.ModelStudent, schemas.SchemaDatabaseError)
}

type serviceResults struct {
	repository repositorys.RepositoryResults
}

func NewServiceResults(repository repositorys.RepositoryResults) *serviceResults {
	return &serviceResults{repository: repository}
}

func (s *serviceResults) ResultsStudentService() (*[]model.ModelStudent, schemas.SchemaDatabaseError) {

	res, err := s.repository.ResultsStudentRepository()
	return res, err
}
