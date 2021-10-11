package repositorys

import (
	"net/http"

	"gorm.io/gorm"

	"github.com/restuwahyu13/gin-rest-api/models"
	"github.com/restuwahyu13/gin-rest-api/schemas"
)

type RepositoryRegister interface {
	RegisterRepository(input *schemas.SchemaAuth) (*models.ModelAuth, schemas.SchemaDatabaseError)
}

type repositoryRegister struct {
	db *gorm.DB
}

func NewRepositoryRegister(db *gorm.DB) *repositoryRegister {
	return &repositoryRegister{db: db}
}

func (r *repositoryRegister) RegisterRepository(input *schemas.SchemaAuth) (*models.ModelAuth, schemas.SchemaDatabaseError) {

	var user models.ModelAuth
	db := r.db.Model(&user)
	errorCode := make(chan schemas.SchemaDatabaseError, 1)

	checkUserAccount := db.Debug().First(&user, "email = ?", input.Email)

	if checkUserAccount.RowsAffected > 0 {
		errorCode <- schemas.SchemaDatabaseError{
			Code: http.StatusConflict,
			Type: "error_01",
		}
		return &user, <-errorCode
	}

	user.Fullname = input.Fullname
	user.Email = input.Email
	user.Password = input.Password

	addNewUser := db.Debug().Create(&user).Commit()

	if addNewUser.RowsAffected < 1 {
		errorCode <- schemas.SchemaDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_02",
		}
		return &user, <-errorCode
	}

	return &user, <-errorCode
}
