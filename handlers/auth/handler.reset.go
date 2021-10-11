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

type handlerReset struct {
	service services.ServiceReset
}

func NewHandlerReset(service services.ServiceReset) *handlerReset {
	return &handlerReset{service: service}
}

func (h *handlerReset) ResetHandler(ctx *gin.Context) {

	var input schemas.SchemaAuth
	ctx.ShouldBindJSON(&input)

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
				Field:   "Password",
				Message: "password is required on body",
			},
			gpc.ErrorMetaConfig{
				Tag:     "gte",
				Field:   "Password",
				Message: "password minimum must be 8 character",
			},
			gpc.ErrorMetaConfig{
				Tag:     "required",
				Field:   "Cpassword",
				Message: "cpassword is required on body",
			},
			gpc.ErrorMetaConfig{
				Tag:     "gte",
				Field:   "Cpassword",
				Message: "cpassword minimum must be 8 character",
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

	if input.Cpassword != input.Password {
		helpers.APIResponse(ctx, "Confirm Password is not match with Password", http.StatusBadRequest, http.MethodPost, nil)
		return
	}

	result := pkg.DecodeToken(resultToken)
	input.Email = result.Claims.Email
	input.Active = true

	_, err := h.service.ResetService(&input)

	switch err.Type {
	case "error_01":
		helpers.APIResponse(ctx, "User account is not exist", err.Code, http.MethodPost, nil)
		return
	case "error_02":
		helpers.APIResponse(ctx, "User account is not active", err.Code, http.MethodPost, nil)
		return
	case "error_03":
		helpers.APIResponse(ctx, "Change new password failed", err.Code, http.MethodPost, nil)
		return
	default:
		helpers.APIResponse(ctx, "Change new password successfully", http.StatusOK, http.MethodPost, nil)
	}
}
