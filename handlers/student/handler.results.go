package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/restuwahyu13/gin-rest-api/helpers"
	services "github.com/restuwahyu13/gin-rest-api/services/student"
)

type handlerResults struct {
	service services.ServiceResults
}

func NewHandlerResultsStudent(service services.ServiceResults) *handlerResults {
	return &handlerResults{service: service}
}

func (h *handlerResults) ResultsStudentHandler(ctx *gin.Context) {

	res, err := h.service.ResultsStudentService()

	switch err.Type {
	case "error_01":
		helpers.APIResponse(ctx, "Students data is not exists", err.Code, http.MethodPost, nil)
	default:
		helpers.APIResponse(ctx, "Results Students data successfully", http.StatusOK, http.MethodPost, res)
	}
}
