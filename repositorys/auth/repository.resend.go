package repositorys

import (
	"net/http"

	"gorm.io/gorm"

	"github.com/restuwahyu13/gin-rest-api/models"
	"github.com/restuwahyu13/gin-rest-api/schemas"
)

type RepositoryResend interface {
	ResendRepository(input *schemas.SchemaAuth) (*models.ModelAuth, schemas.SchemaDatabaseError)
}

type repositoryResend struct {
	db *gorm.DB
}

func NewRepositoryResend(db *gorm.DB) *repositoryResend {
	return &repositoryResend{db: db}
}

func (r *repositoryResend) ResendRepository(input *schemas.SchemaAuth) (*models.ModelAuth, schemas.SchemaDatabaseError) {

	var user models.ModelAuth
	db := r.db.Model(&user)
	errorCode := make(chan schemas.SchemaDatabaseError, 1)

	user.Email = input.Email

	checkUserAccount := db.Debug().First(&user, "email = ?", input.Email)

	if checkUserAccount.RowsAffected < 1 {
		errorCode <- schemas.SchemaDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_01",
		}
		return &user, <-errorCode
	}

	if user.Active {
		errorCode <- schemas.SchemaDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_02",
		}
		return &user, <-errorCode
	}

	return &user, <-errorCode
}
