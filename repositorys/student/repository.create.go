package repositorys

import (
	"net/http"

	"gorm.io/gorm"

	"github.com/restuwahyu13/gin-rest-api/models"
	"github.com/restuwahyu13/gin-rest-api/schemas"
)

type RepositoryCreate interface {
	CreateStudentRepository(input *schemas.SchemaStudent) (*models.ModelStudent, schemas.SchemaDatabaseError)
}

type repositoryCreate struct {
	db *gorm.DB
}

func NewRepositoryCreate(db *gorm.DB) *repositoryCreate {
	return &repositoryCreate{db: db}
}

func (r *repositoryCreate) CreateStudentRepository(input *schemas.SchemaStudent) (*models.ModelStudent, schemas.SchemaDatabaseError) {

	var students models.ModelStudent
	db := r.db.Model(&students)
	errorCode := make(chan schemas.SchemaDatabaseError, 1)

	checkStudentExist := db.Debug().First(&students, "npm = ?", input.Npm)

	if checkStudentExist.RowsAffected > 0 {
		errorCode <- schemas.SchemaDatabaseError{
			Code: http.StatusConflict,
			Type: "error_01",
		}
		return &students, <-errorCode
	}

	students.Name = input.Name
	students.Npm = input.Npm
	students.Fak = input.Fak
	students.Bid = input.Bid

	addNewStudent := db.Debug().Create(&students).Commit()

	if addNewStudent.RowsAffected < 1 {
		errorCode <- schemas.SchemaDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_02",
		}
		return &students, <-errorCode
	}

	return &students, <-errorCode
}
