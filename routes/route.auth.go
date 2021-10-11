package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	handlers "github.com/restuwahyu13/gin-rest-api/handlers/auth"
	repositorys "github.com/restuwahyu13/gin-rest-api/repositorys/auth"
	services "github.com/restuwahyu13/gin-rest-api/services/auth"
)

func InitAuthRoutes(db *gorm.DB, route *gin.Engine) {

	/**
	@description All Handler Auth
	*/

	pingRepository := repositorys.NewRepositoryPing(db)
	pingService := services.NewServicePing(pingRepository)
	pingHandler := handlers.NewHandlerPing(pingService)

	LoginRepository := repositorys.NewRepositoryLogin(db)
	loginService := services.NewServiceLogin(LoginRepository)
	loginHandler := handlers.NewHandlerLogin(loginService)

	registerRepository := repositorys.NewRepositoryRegister(db)
	registerService := services.NewServiceRegister(registerRepository)
	registerHandler := handlers.NewHandlerRegister(registerService)

	activationRepository := repositorys.NewRepositoryActivation(db)
	activationService := services.NewServiceActivation(activationRepository)
	activationHandler := handlers.NewHandlerActivation(activationService)

	resendRepository := repositorys.NewRepositoryResend(db)
	resendService := services.NewServiceResend(resendRepository)
	resendHandler := handlers.NewHandlerResend(resendService)

	forgotRepository := repositorys.NewRepositoryForgot(db)
	forgotService := services.NewServiceForgot(forgotRepository)
	forgotHandler := handlers.NewHandlerForgot(forgotService)

	resetRepository := repositorys.NewRepositoryReset(db)
	resetService := services.NewServiceReset(resetRepository)
	resetHandler := handlers.NewHandlerReset(resetService)

	/**
	@description All Auth Route
	*/
	groupRoute := route.Group("/api/v1")
	groupRoute.GET("/ping", pingHandler.PingHandler)
	groupRoute.POST("/register", registerHandler.RegisterHandler)
	groupRoute.POST("/login", loginHandler.LoginHandler)
	groupRoute.POST("/activation/:token", activationHandler.ActivationHandler)
	groupRoute.POST("/resend-token", resendHandler.ResendHandler)
	groupRoute.POST("/forgot-password", forgotHandler.ForgotHandler)
	groupRoute.POST("/change-password/:token", resetHandler.ResetHandler)

}
