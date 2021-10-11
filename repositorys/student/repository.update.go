package repositorys

import (
	"net/http"

	"gorm.io/gorm"

	"github.com/restuwahyu13/gin-rest-api/models"
	"github.com/restuwahyu13/gin-rest-api/schemas"
)

type RepositoryUpdate interface {
	UpdateStudentRepository(input *schemas.SchemaStudent) (*models.ModelStudent, schemas.SchemaDatabaseError)
}

type repositoryUpdate struct {
	db *gorm.DB
}

func NewRepositoryUpdate(db *gorm.DB) *repositoryUpdate {
	return &repositoryUpdate{db: db}
}

func (r *repositoryUpdate) UpdateStudentRepository(input *schemas.SchemaStudent) (*models.ModelStudent, schemas.SchemaDatabaseError) {

	var students models.ModelStudent
	db := r.db.Model(&students)
	errorCode := make(chan schemas.SchemaDatabaseError, 1)

	students.ID = input.ID

	checkStudentId := db.Debug().First(&students)

	if checkStudentId.RowsAffected < 1 {
		errorCode <- schemas.SchemaDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_01",
		}
		return &students, <-errorCode
	}

	students.Name = input.Name
	students.Npm = input.Npm
	students.Fak = input.Fak
	students.Bid = input.Bid

	updateStudent := db.Debug().Where("id = ?", input.ID).Updates(students)

	if updateStudent.RowsAffected < 1 {
		errorCode <- schemas.SchemaDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_02",
		}
		return &students, <-errorCode
	}

	return &students, <-errorCode
}
