package repositorys

import (
	"net/http"

	"gorm.io/gorm"

	"github.com/restuwahyu13/gin-rest-api/models"
	"github.com/restuwahyu13/gin-rest-api/pkg"
	"github.com/restuwahyu13/gin-rest-api/schemas"
)

type RepositoryLogin interface {
	LoginRepository(input *schemas.SchemaAuth) (*models.ModelAuth, schemas.SchemaDatabaseError)
}

type repositoryLogin struct {
	db *gorm.DB
}

func NewRepositoryLogin(db *gorm.DB) *repositoryLogin {
	return &repositoryLogin{db: db}
}

func (r *repositoryLogin) LoginRepository(input *schemas.SchemaAuth) (*models.ModelAuth, schemas.SchemaDatabaseError) {

	var user models.ModelAuth
	db := r.db.Model(&user)
	errorCode := make(chan schemas.SchemaDatabaseError, 1)

	user.Email = input.Email
	user.Password = input.Password

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

	comparePassword := pkg.ComparePassword(user.Password, input.Password)

	if comparePassword != nil {
		errorCode <- schemas.SchemaDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_03",
		}
		return &user, <-errorCode
	}

	return &user, <-errorCode
}
