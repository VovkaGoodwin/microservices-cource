package http_handler

import "github.com/gin-gonic/gin"

func (h *Handler) initRoutes() *gin.Engine {
	r := gin.New()

	r.GET("/liveness", h.liveness)

	return r
}
