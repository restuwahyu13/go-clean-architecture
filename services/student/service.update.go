package student

import (
	"github.com/restuwahyu13/gin-rest-api/models"
	repositorys "github.com/restuwahyu13/gin-rest-api/repositorys/student"
	"github.com/restuwahyu13/gin-rest-api/schemas"
)

type ServiceUpdate interface {
	UpdateStudentService(input *schemas.SchemaStudent) (*models.ModelStudent, schemas.SchemaDatabaseError)
}

type serviceUpdate struct {
	repository repositorys.RepositoryUpdate
}

func NewServiceUpdate(repository repositorys.RepositoryUpdate) *serviceUpdate {
	return &serviceUpdate{repository: repository}
}

func (s *serviceUpdate) UpdateStudentService(input *schemas.SchemaStudent) (*models.ModelStudent, schemas.SchemaDatabaseError) {

	var student schemas.SchemaStudent
	student.ID = input.ID
	student.Name = input.Name
	student.Npm = input.Npm
	student.Fak = input.Fak
	student.Bid = input.Fak

	res, err := s.repository.UpdateStudentRepository(&student)
	return res, err
}
