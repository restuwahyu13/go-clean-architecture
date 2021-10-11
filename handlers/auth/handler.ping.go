package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/restuwahyu13/gin-rest-api/helpers"
	services "github.com/restuwahyu13/gin-rest-api/services/auth"
)

type handlerPing struct {
	service services.ServicePing
}

func NewHandlerPing(service services.ServicePing) *handlerPing {
	return &handlerPing{service: service}
}

func (h *handlerPing) PingHandler(ctx *gin.Context) {
	res := h.service.PingService()
	helpers.APIResponse(ctx, res, http.StatusOK, http.MethodPost, nil)
}
