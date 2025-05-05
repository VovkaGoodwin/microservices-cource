package http_server

import (
	"github.com/VovkaGoodwin/microservices-cource/pkg/config"
	"github.com/VovkaGoodwin/microservices-cource/pkg/healthcheck"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	cfg *config.Config
	hc  healthcheck.Interactor
}

func NewHandler(cfg *config.Config, healthCheck healthcheck.Interactor) *gin.Engine {
	handler := Handler{
		cfg: cfg,
		hc:  healthCheck,
	}

	return handler.InitRoutes()
}

func (h *Handler) respondError(c *gin.Context, statusCode int, message interface{}) {
	c.AbortWithStatusJSON(statusCode, gin.H{
		"message": message,
	})
}

func (h *Handler) respondSuccess(c *gin.Context, statusCode int, data interface{}) {
	c.JSON(statusCode, gin.H{
		"data": data,
	})
}
