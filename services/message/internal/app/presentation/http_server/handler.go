package http_server

import (
	"github.com/VovkaGoodwin/microservices-cource/services/message/internal/app/application/interactors/healthcheck"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	hc healthcheck.Interactor
}

func NewHttpHandler(hc healthcheck.Interactor) *gin.Engine {
	handler := &Handler{hc}

	return handler.InitRoutes()
}

func (h *Handler) respondSuccess(c *gin.Context, statusCode int, data interface{}) {
	c.JSON(statusCode, gin.H{
		"data": data,
	})
}
