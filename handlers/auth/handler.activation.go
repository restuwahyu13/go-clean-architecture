package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	gpc "github.com/restuwahyu13/go-playground-converter"
	"github.com/sirupsen/logrus"

	"github.com/restuwahyu13/gin-rest-api/helpers"
	"github.com/restuwahyu13/gin-rest-api/pkg"
	"github.com/restuwahyu13/gin-rest-api/schemas"
	services "github.com/restuwahyu13/gin-rest-api/services/auth"
)

type handlerActivation struct {
	service services.ServiceActivation
}

func NewHandlerActivation(service services.ServiceActivation) *handlerActivation {
	return &handlerActivation{service: service}
}

func (h *handlerActivation) ActivationHandler(ctx *gin.Context) {

	var input schemas.SchemaAuth

	config := gpc.ErrorConfig{
		Options: []gpc.ErrorMetaConfig{
			gpc.ErrorMetaConfig{
				Tag:     "required",
				Field:   "Email",
				Message: "email is required on body",
			},
			gpc.ErrorMetaConfig{
				Tag:     "email",
				Field:   "Email",
				Message: "email format is not valid",
			},
			gpc.ErrorMetaConfig{
				Tag:     "required",
				Field:   "Token",
				Message: "accessToken is required on params",
			},
		},
	}

	errResponse, errCount := pkg.GoValidator(input, config.Options)

	if errCount > 0 {
		helpers.ValidatorErrorResponse(ctx, http.StatusBadRequest, http.MethodPost, errResponse)
		return
	}

	token := ctx.Param("token")
	resultToken, errToken := pkg.VerifyToken(token, "JWT_SECRET")

	if errToken != nil {
		defer logrus.Error(errToken.Error())
		helpers.APIResponse(ctx, "Verified activation token failed", http.StatusBadRequest, http.MethodPost, nil)
		return
	}

	result := pkg.DecodeToken(resultToken)
	input.Email = result.Claims.Email
	input.Active = true

	_, err := h.service.ActivationService(&input)

	switch err.Type {
	case "error_01":
		helpers.APIResponse(ctx, "User account is not exist", err.Code, http.MethodPost, nil)
		return
	case "error_02":
		helpers.APIResponse(ctx, "User account hash been active please login", err.Code, http.MethodPost, nil)
		return
	case "error_03":
		helpers.APIResponse(ctx, "Activation account failed", err.Code, http.MethodPost, nil)
		return
	default:
		helpers.APIResponse(ctx, "Activation account success", http.StatusOK, http.MethodPost, nil)
	}
}
