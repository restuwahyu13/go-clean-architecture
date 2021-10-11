package repositorys

import (
	"net/http"

	"gorm.io/gorm"

	"github.com/restuwahyu13/gin-rest-api/models"
	"github.com/restuwahyu13/gin-rest-api/schemas"
)

type RepositoryResults interface {
	ResultsStudentRepository() (*[]models.ModelStudent, schemas.SchemaDatabaseError)
}

type repositoryResults struct {
	db *gorm.DB
}

func NewRepositoryResults(db *gorm.DB) *repositoryResults {
	return &repositoryResults{db: db}
}

func (r *repositoryResults) ResultsStudentRepository() (*[]models.ModelStudent, schemas.SchemaDatabaseError) {

	var students []models.ModelStudent
	db := r.db.Model(&students)
	errorCode := make(chan schemas.SchemaDatabaseError, 1)

	resultsStudents := db.Debug().Find(&students)

	if resultsStudents.RowsAffected < 1 {
		errorCode <- schemas.SchemaDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_01",
		}
		return &students, <-errorCode
	}

	return &students, <-errorCode
}
