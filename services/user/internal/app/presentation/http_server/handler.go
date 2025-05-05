package http_server

import (
	"github.com/VovkaGoodwin/microservices-cource/services/user/internal/app/application/interactors/healthcheck"
	"github.com/gin-gonic/gin"
)

type Deps struct {
	hc healthcheck.Interactor
}

type Handler struct {
	*Deps
}

func NewHttpHandler(deps *Deps) *gin.Engine {
	handler := &Handler{
		deps,
	}

	return handler.initRoutes()
}

func (h *Handler) respondSuccess(c *gin.Context, statusCode int, data interface{}) {
	c.JSON(statusCode, gin.H{
		"data": data,
	})
}
