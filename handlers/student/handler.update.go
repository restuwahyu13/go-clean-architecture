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

type handlerUpdate struct {
	service services.ServiceUpdate
}

func NewHandlerUpdateStudent(service services.ServiceUpdate) *handlerUpdate {
	return &handlerUpdate{service: service}
}

func (h *handlerUpdate) UpdateStudentHandler(ctx *gin.Context) {

	var input schemas.SchemaStudent
	input.ID = ctx.Param("id")
	ctx.ShouldBindJSON(&input)

	config := gpc.ErrorConfig{
		Options: []gpc.ErrorMetaConfig{
			gpc.ErrorMetaConfig{
				Tag:     "required",
				Field:   "ID",
				Message: "id is required on param",
			},
			gpc.ErrorMetaConfig{
				Tag:     "uuid",
				Field:   "ID",
				Message: "params must be uuid format",
			},
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
				Tag:     "number",
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
		helpers.ValidatorErrorResponse(ctx, http.StatusBadRequest, http.MethodGet, errResponse)
		return
	}

	_, err := h.service.UpdateStudentService(&input)

	switch err.Type {
	case "error_01":
		helpers.APIResponse(ctx, "Student data is not exist or deleted", http.StatusNotFound, http.MethodPost, nil)
	case "error_02":
		helpers.APIResponse(ctx, "Update student data failed", http.StatusForbidden, http.MethodPost, nil)
	default:
		helpers.APIResponse(ctx, "Update student data sucessfully", http.StatusOK, http.MethodPost, nil)
	}
}
