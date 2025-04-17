package http_handler

import (
	"auth/internal/app/application/interactors"
	"auth/internal/config"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	cfg *config.Config
	hc  *interactors.Healthcheck
}

func NewHandler(cfg *config.Config, healthCheck *interactors.Healthcheck) *Handler {
	return &Handler{
		cfg: cfg,
		hc:  healthCheck,
	}
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
