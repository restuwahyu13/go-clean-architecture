package repositorys

import (
	"net/http"

	"gorm.io/gorm"

	"github.com/restuwahyu13/gin-rest-api/models"
	"github.com/restuwahyu13/gin-rest-api/schemas"
)

type RepositoryDelete interface {
	DeleteStudentRepository(input *schemas.SchemaStudent) (*models.ModelStudent, schemas.SchemaDatabaseError)
}

type repositoryDelete struct {
	db *gorm.DB
}

func NewRepositoryDelete(db *gorm.DB) *repositoryDelete {
	return &repositoryDelete{db: db}
}

func (r *repositoryDelete) DeleteStudentRepository(input *schemas.SchemaStudent) (*models.ModelStudent, schemas.SchemaDatabaseError) {

	var students models.ModelStudent
	db := r.db.Model(&students)
	errorCode := make(chan schemas.SchemaDatabaseError, 1)

	checkStudentId := db.Debug().First(&students)

	if checkStudentId.RowsAffected < 1 {
		errorCode <- schemas.SchemaDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_01",
		}
		return &students, <-errorCode
	}

	deleteStudentId := db.Debug().Delete(&students)

	if deleteStudentId.RowsAffected < 1 {
		errorCode <- schemas.SchemaDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_02",
		}
		return &students, <-errorCode
	}

	return &students, <-errorCode
}
