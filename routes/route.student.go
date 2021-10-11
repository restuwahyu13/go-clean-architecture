package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	handlers "github.com/restuwahyu13/gin-rest-api/handlers/student"
	middleware "github.com/restuwahyu13/gin-rest-api/middlewares"
	repositorys "github.com/restuwahyu13/gin-rest-api/repositorys/student"
	services "github.com/restuwahyu13/gin-rest-api/services/student"
)

func InitStudentRoutes(db *gorm.DB, route *gin.Engine) {

	/**
	@description All Handler Student
	*/
	createStudentRepository := repositorys.NewRepositoryCreate(db)
	createStudentService := services.NewServiceCreate(createStudentRepository)
	createStudentHandler := handlers.NewHandlerCreateStudent(createStudentService)

	resultsStudentRepository := repositorys.NewRepositoryResults(db)
	resultsStudentService := services.NewServiceResults(resultsStudentRepository)
	resultsStudentHandler := handlers.NewHandlerResultsStudent(resultsStudentService)

	resultStudentRepository := repositorys.NewRepositoryResult(db)
	resultStudentService := services.NewServiceResult(resultStudentRepository)
	resultStudentHandler := handlers.NewHandlerResultStudent(resultStudentService)

	deleteStudentRepository := repositorys.NewRepositoryDelete(db)
	deleteStudentService := services.NewServiceDelete(deleteStudentRepository)
	deleteStudentHandler := handlers.NewHandlerDeleteStudent(deleteStudentService)

	updateStudentRepository := repositorys.NewRepositoryUpdate(db)
	updateStudentService := services.NewServiceUpdate(updateStudentRepository)
	updateStudentHandler := handlers.NewHandlerUpdateStudent(updateStudentService)

	/**
	@description All Student Route
	*/
	groupRoute := route.Group("/api/v1").Use(middleware.Auth())
	groupRoute.POST("/student", createStudentHandler.CreateStudentHandler)
	groupRoute.GET("/student", resultsStudentHandler.ResultsStudentHandler)
	groupRoute.GET("/student/:id", resultStudentHandler.ResultStudentHandler)
	groupRoute.DELETE("/student/:id", deleteStudentHandler.DeleteStudentHandler)
	groupRoute.PUT("/student/:id", updateStudentHandler.UpdateStudentHandler)
}
