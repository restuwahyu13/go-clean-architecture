package repositorys

import (
	"net/http"

	"gorm.io/gorm"

	"github.com/restuwahyu13/gin-rest-api/models"
	"github.com/restuwahyu13/gin-rest-api/schemas"
)

type RepositoryResult interface {
	ResultStudentRepository(input *schemas.SchemaStudent) (*models.ModelStudent, schemas.SchemaDatabaseError)
}

type repositoryResult struct {
	db *gorm.DB
}

func NewRepositoryResult(db *gorm.DB) *repositoryResult {
	return &repositoryResult{db: db}
}

func (r *repositoryResult) ResultStudentRepository(input *schemas.SchemaStudent) (*models.ModelStudent, schemas.SchemaDatabaseError) {

	var students models.ModelStudent
	db := r.db.Model(&students)
	errorCode := make(chan schemas.SchemaDatabaseError, 1)

	resultStudents := db.Debug().First(&students)

	if resultStudents.RowsAffected < 1 {
		errorCode <- schemas.SchemaDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_01",
		}
		return &students, <-errorCode
	}

	return &students, <-errorCode
}
