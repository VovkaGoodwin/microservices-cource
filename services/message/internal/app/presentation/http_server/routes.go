package http_server

import (
	"github.com/gin-gonic/gin"
)

func (h *Handler) InitRoutes() *gin.Engine {
	r := gin.New()
	r.GET("/liveness", h.liveness)

	return r
}
