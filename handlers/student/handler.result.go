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

type handlerResult struct {
	service services.ServiceResult
}

func NewHandlerResultStudent(service services.ServiceResult) *handlerResult {
	return &handlerResult{service: service}
}

func (h *handlerResult) ResultStudentHandler(ctx *gin.Context) {

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
		helpers.ValidatorErrorResponse(ctx, http.StatusBadRequest, http.MethodGet, errResponse)
		return
	}

	res, err := h.service.ResultStudentService(&input)

	switch err.Type {
	case "error_01":
		helpers.APIResponse(ctx, "Student data is not exist or deleted", err.Code, http.MethodGet, nil)
		return
	default:
		helpers.APIResponse(ctx, "Result Student data successfully", http.StatusOK, http.MethodGet, res)
	}
}
