package configs

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/restuwahyu13/gin-rest-api/models"
	"github.com/restuwahyu13/gin-rest-api/routes"
)

type Testing interface {
	SetupRouter() *gin.Engine
	SetupDatabase() *gorm.DB
}

type RouterTesting struct {
	testing Testing
}

func NewRouterTesting() *gin.Engine {
	var router RouterTesting
	return router.SetupRouter()
}

func (r *RouterTesting) SetupRouter() *gin.Engine {
	db := r.SetupDatabase()
	app := gin.Default()

	gin.SetMode(gin.TestMode)

	routes.InitAuthRoutes(db, app)
	routes.InitStudentRoutes(db, app)

	return app
}

func (r *RouterTesting) SetupDatabase() *gorm.DB {
	db, err := gorm.Open(postgres.Open("postgres://agtran:agtran@localhost:5432/agtran"), &gorm.Config{})

	if err != nil {
		defer logrus.Info("Connect into Database Failed")
		logrus.Fatal(err.Error())
	}

	err = db.AutoMigrate(
		&models.ModelAuth{},
		&models.ModelStudent{},
	)

	if err != nil {
		logrus.Fatal(err.Error())
	}

	return db
}
