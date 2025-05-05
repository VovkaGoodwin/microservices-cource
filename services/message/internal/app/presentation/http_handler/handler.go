package http_handler

import (
	"github.com/VovkaGoodwin/microservices-cource/pkg/healthcheck"
	"github.com/gin-gonic/gin"
)

type Deps struct {
	hc healthcheck.Interactor
}

type Handler struct {
	*Deps
}

func NewHttpHandler(deps *Deps) *gin.Engine {
	handler := &Handler{deps}

	return handler.InitRoutes()
}

func (h *Handler) respondSuccess(c *gin.Context, statusCode int, data interface{}) {
	c.JSON(statusCode, gin.H{
		"data": data,
	})
}
