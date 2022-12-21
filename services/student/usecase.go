package student

import (
	repositorys "github.com/evermos/trial/go-clean-architecture/repositorys/student"
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
	repositoryCreate  repositorys.RepositoryCreate
	repositoryDelete  repositorys.RepositoryDelete
	repositoryResult  repositorys.RepositoryResult
	repositoryResults repositorys.RepositoryResults
	repositoryUpdate  repositorys.RepositoryUpdate
}

func NewService(
	repositoryCreate repositorys.RepositoryCreate,
	repositoryDelete repositorys.RepositoryDelete,
	repositoryResult repositorys.RepositoryResult,
	repositoryResults repositorys.RepositoryResults,
	repositoryUpdate repositorys.RepositoryUpdate,
) *service {
	return &service{
		repositoryCreate:  repositoryCreate,
		repositoryDelete:  repositoryDelete,
		repositoryResult:  repositoryResult,
		repositoryResults: repositoryResults,
		repositoryUpdate:  repositoryUpdate,
	}
}
