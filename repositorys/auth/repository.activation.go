package repositorys

import (
	"net/http"
	"time"

	"gorm.io/gorm"

	"github.com/restuwahyu13/gin-rest-api/models"
	"github.com/restuwahyu13/gin-rest-api/schemas"
)

type RepositoryActivation interface {
	ActivationRepository(input *schemas.SchemaAuth) (*models.ModelAuth, schemas.SchemaDatabaseError)
}

type repositoryActivation struct {
	db *gorm.DB
}

func NewRepositoryActivation(db *gorm.DB) *repositoryActivation {
	return &repositoryActivation{db: db}
}

func (r *repositoryActivation) ActivationRepository(input *schemas.SchemaAuth) (*models.ModelAuth, schemas.SchemaDatabaseError) {

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

	db.Debug().First(&user, "activation = ?", input.Active)

	if user.Active {
		errorCode <- schemas.SchemaDatabaseError{
			Code: http.StatusBadRequest,
			Type: "error_02",
		}
		return &user, <-errorCode
	}

	user.Active = input.Active
	user.UpdatedAt = time.Now().Local()

	updateActivation := db.Debug().Where("email = ?", input.Email).Updates(user)

	if updateActivation.RowsAffected < 1 {
		errorCode <- schemas.SchemaDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_03",
		}
		return &user, <-errorCode
	}

	return &user, <-errorCode
}
