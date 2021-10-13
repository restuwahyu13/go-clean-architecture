package student

import (
	"github.com/restuwahyu13/gin-rest-api/models"
	model "github.com/restuwahyu13/gin-rest-api/models"
	"github.com/restuwahyu13/gin-rest-api/schemas"
)

func (s *service) Result(input *schemas.SchemaStudent) (*models.ModelStudent, schemas.SchemaDatabaseError) {

	var student schemas.SchemaStudent
	student.ID = input.ID

	res, err := s.repository.ResultStudentRepository(&student)
	return res, err
}

func (s *service) ResultsStudentService() (*[]model.ModelStudent, schemas.SchemaDatabaseError) {

	res, err := s.repository.ResultsStudentRepository()
	return res, err
}

func (s *service) DeleteStudentService(input *schemas.SchemaStudent) (*models.ModelStudent, schemas.SchemaDatabaseError) {

	var student schemas.SchemaStudent
	student.ID = input.ID

	res, err := s.repository.DeleteStudentRepository(&student)
	return res, err
}

func (s *service) CreateStudentService(input *schemas.SchemaStudent) (*models.ModelStudent, schemas.SchemaDatabaseError) {

	var student schemas.SchemaStudent
	student.Name = input.Name
	student.Npm = input.Npm
	student.Fak = input.Fak
	student.Bid = input.Bid

	res, err := s.repository.CreateStudentRepository(&student)
	return res, err
}

func (s *service) UpdateStudentService(input *schemas.SchemaStudent) (*models.ModelStudent, schemas.SchemaDatabaseError) {

	var student schemas.SchemaStudent
	student.ID = input.ID
	student.Name = input.Name
	student.Npm = input.Npm
	student.Fak = input.Fak
	student.Bid = input.Fak

	res, err := s.repository.UpdateStudentRepository(&student)
	return res, err
}
