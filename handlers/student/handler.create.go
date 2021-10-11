package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	gpc "github.com/restuwahyu13/go-playground-converter"

	"github.com/restuwahyu13/gin-rest-api/helpers"
	"github.com/restuwahyu13/gin-rest-api/pkg"
	"github.com/restuwahyu13/gin-rest-api/schemas"
	services "github.com/restuwahyu13/gin-rest-api/services/student"
)

type handlerCreate struct {
	service services.ServiceCreate
}

func NewHandlerCreateStudent(service services.ServiceCreate) *handlerCreate {
	return &handlerCreate{service: service}
}

func (h *handlerCreate) CreateStudentHandler(ctx *gin.Context) {

	var input schemas.SchemaStudent
	ctx.ShouldBindJSON(&input)

	config := gpc.ErrorConfig{
		Options: []gpc.ErrorMetaConfig{
			gpc.ErrorMetaConfig{
				Tag:     "required",
				Field:   "Name",
				Message: "name is required on body",
			},
			gpc.ErrorMetaConfig{
				Tag:     "lowercase",
				Field:   "Name",
				Message: "name must be using lowercase",
			},
			gpc.ErrorMetaConfig{
				Tag:     "required",
				Field:   "Npm",
				Message: "npm is required on body",
			},
			gpc.ErrorMetaConfig{
				Tag:     "numeric",
				Field:   "Npm",
				Message: "npm must be number format",
			},
			gpc.ErrorMetaConfig{
				Tag:     "required",
				Field:   "Fak",
				Message: "fak is required on body",
			},
			gpc.ErrorMetaConfig{
				Tag:     "lowercase",
				Field:   "Fak",
				Message: "fak must be using lowercase",
			},
			gpc.ErrorMetaConfig{
				Tag:     "required",
				Field:   "Bid",
				Message: "bid is required on body",
			},
			gpc.ErrorMetaConfig{
				Tag:     "lowercase",
				Field:   "Bid",
				Message: "bid must be using lowercase",
			},
		},
	}

	errResponse, errCount := pkg.GoValidator(&input, config.Options)

	if errCount > 0 {
		helpers.ValidatorErrorResponse(ctx, http.StatusBadRequest, http.MethodPost, errResponse)
		return
	}

	_, err := h.service.CreateStudentService(&input)

	switch err.Type {
	case "error_01":
		helpers.APIResponse(ctx, "Npm student already exist", err.Code, http.MethodPost, nil)
		return
	case "error_02":
		helpers.APIResponse(ctx, "Create new student account failed", err.Code, http.MethodPost, nil)
		return
	default:
		helpers.APIResponse(ctx, "Create new student account successfully", http.StatusCreated, http.MethodPost, nil)
	}
}
