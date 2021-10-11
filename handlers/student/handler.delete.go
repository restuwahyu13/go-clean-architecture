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

type handlerDelete struct {
	service services.ServiceDelete
}

func NewHandlerDeleteStudent(service services.ServiceDelete) *handlerDelete {
	return &handlerDelete{service: service}
}

func (h *handlerDelete) DeleteStudentHandler(ctx *gin.Context) {

	var input schemas.SchemaStudent
	input.ID = ctx.Param("id")

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
		},
	}

	errResponse, errCount := pkg.GoValidator(&input, config.Options)

	if errCount > 0 {
		helpers.ValidatorErrorResponse(ctx, http.StatusBadRequest, http.MethodDelete, errResponse)
		return
	}

	_, err := h.service.DeleteStudentService(&input)

	switch err.Type {
	case "error_01":
		helpers.APIResponse(ctx, "Student data is not exist or deleted", err.Code, http.MethodDelete, nil)
		return
	case "error_02":
		helpers.APIResponse(ctx, "Delete student data failed", err.Code, http.MethodDelete, nil)
		return
	default:
		helpers.APIResponse(ctx, "Delete student data successfully", http.StatusOK, http.MethodDelete, nil)
	}
}
