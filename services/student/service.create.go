package student

import (
	"github.com/restuwahyu13/gin-rest-api/models"
	repositorys "github.com/restuwahyu13/gin-rest-api/repositorys/student"
	"github.com/restuwahyu13/gin-rest-api/schemas"
)

type ServiceCreate interface {
	CreateStudentService(input *schemas.SchemaStudent) (*models.ModelStudent, schemas.SchemaDatabaseError)
}

type serviceCreate struct {
	repository repositorys.RepositoryCreate
}

func NewServiceCreate(repository repositorys.RepositoryCreate) *serviceCreate {
	return &serviceCreate{repository: repository}
}

func (s *serviceCreate) CreateStudentService(input *schemas.SchemaStudent) (*models.ModelStudent, schemas.SchemaDatabaseError) {

	var student schemas.SchemaStudent
	student.Name = input.Name
	student.Npm = input.Npm
	student.Fak = input.Fak
	student.Bid = input.Bid

	res, err := s.repository.CreateStudentRepository(&student)
	return res, err
}
