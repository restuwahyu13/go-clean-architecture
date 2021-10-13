package student

import (
	"github.com/restuwahyu13/gin-rest-api/models"
	repositorys "github.com/restuwahyu13/gin-rest-api/repositorys/student"
	"github.com/restuwahyu13/gin-rest-api/schemas"
)

type Service interface {
	Create(input *schemas.SchemaStudent) (*models.ModelStudent, schemas.SchemaDatabaseError)
	Delete(input *schemas.SchemaStudent) (*models.ModelStudent, schemas.SchemaDatabaseError)
	Result(input *schemas.SchemaStudent) (*models.ModelStudent, schemas.SchemaDatabaseError)
	Results() (*[]models.ModelStudent, schemas.SchemaDatabaseError)
	Update(input *schemas.SchemaStudent) (*models.ModelStudent, schemas.SchemaDatabaseError)
}

type service struct {
	repository repositorys.RepositoryCreate
}

func NewService(repository repositorys.RepositoryCreate) *serviceCreate {
	return &serviceCreate{repository: repository}
}
