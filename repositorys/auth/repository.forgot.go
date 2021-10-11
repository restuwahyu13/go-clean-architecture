package repositorys

import (
	"net/http"

	"gorm.io/gorm"

	"github.com/restuwahyu13/gin-rest-api/models"
	"github.com/restuwahyu13/gin-rest-api/schemas"
)

type RepositoryForgot interface {
	ForgotRepository(input *schemas.SchemaAuth) (*models.ModelAuth, schemas.SchemaDatabaseError)
}

type repositoryForgot struct {
	db *gorm.DB
}

func NewRepositoryForgot(db *gorm.DB) *repositoryForgot {
	return &repositoryForgot{db: db}
}

func (r *repositoryForgot) ForgotRepository(input *schemas.SchemaAuth) (*models.ModelAuth, schemas.SchemaDatabaseError) {

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

	if !user.Active {
		errorCode <- schemas.SchemaDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_02",
		}
		return &user, <-errorCode
	}

	changePassword := db.Debug().Where("email = ?", input.Email).Updates(user)

	if changePassword.RowsAffected < 1 {
		errorCode <- schemas.SchemaDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_03",
		}
		return &user, <-errorCode
	}

	return &user, <-errorCode
}
