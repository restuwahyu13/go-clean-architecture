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

type handlerForgot struct {
	service services.ServiceForgot
}

func NewHandlerForgot(service services.ServiceForgot) *handlerForgot {
	return &handlerForgot{service: service}
}

func (h *handlerForgot) ForgotHandler(ctx *gin.Context) {

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
		},
	}

	errResponse, errCount := pkg.GoValidator(input, config.Options)

	if errCount > 0 {
		helpers.ValidatorErrorResponse(ctx, http.StatusBadRequest, http.MethodPost, errResponse)
		return
	}

	res, err := h.service.ForgotService(&input)

	switch err.Type {
	case "error_01":
		helpers.APIResponse(ctx, "Email is not never registered", err.Code, http.MethodPost, nil)
		return
	case "error_02":
		helpers.APIResponse(ctx, "User account is not active", err.Code, http.MethodPost, nil)
		return
	case "error_03":
		helpers.APIResponse(ctx, "Forgot password failed", err.Code, http.MethodPost, nil)
		return
	default:
		accessTokenData := map[string]interface{}{"id": res.ID, "email": res.Email}
		accessToken, errToken := pkg.Sign(accessTokenData, "JWT_SECRET", 5)

		if errToken != nil {
			defer logrus.Error(errToken.Error())
			helpers.APIResponse(ctx, "Generate accessToken failed", http.StatusBadRequest, http.MethodPost, nil)
			return
		}

		_, errorEmail := pkg.SendGridMail(res.Fullname, res.Email, "Reset Password", "template_reset", accessToken)

		if errorEmail != nil {
			defer logrus.Error(errorEmail.Error())
			helpers.APIResponse(ctx, "Sending email reset password failed", http.StatusBadRequest, http.MethodPost, nil)
			return
		}

		helpers.APIResponse(ctx, "Forgot password successfully", http.StatusOK, http.MethodPost, nil)
	}
}
