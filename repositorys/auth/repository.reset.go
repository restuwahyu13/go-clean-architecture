package repositorys

import (
	"net/http"

	"gorm.io/gorm"

	"github.com/restuwahyu13/gin-rest-api/models"
	"github.com/restuwahyu13/gin-rest-api/schemas"
)

type RepositoryReset interface {
	ResetRepository(input *schemas.SchemaAuth) (*models.ModelAuth, schemas.SchemaDatabaseError)
}

type repositoryReset struct {
	db *gorm.DB
}

func NewRepositoryReset(db *gorm.DB) *repositoryReset {
	return &repositoryReset{db: db}
}

func (r *repositoryReset) ResetRepository(input *schemas.SchemaAuth) (*models.ModelAuth, schemas.SchemaDatabaseError) {

	var users models.ModelAuth
	db := r.db.Model(&users)
	errorCode := make(chan schemas.SchemaDatabaseError, 1)

	users.Email = input.Email
	users.Password = input.Password
	users.Active = input.Active

	checkUserAccount := db.Debug().First(&users, "email = ?", input.Email)

	if checkUserAccount.RowsAffected < 1 {
		errorCode <- schemas.SchemaDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_01",
		}
		return &users, <-errorCode
	}

	if !users.Active {
		errorCode <- schemas.SchemaDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_02",
		}
		return &users, <-errorCode
	}

	updateNewPassword := db.Debug().Where("email = ?", input.Email).Updates(users)

	if updateNewPassword.RowsAffected < 1 {
		errorCode <- schemas.SchemaDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_03",
		}
		return &users, <-errorCode
	}

	return &users, <-errorCode
}
