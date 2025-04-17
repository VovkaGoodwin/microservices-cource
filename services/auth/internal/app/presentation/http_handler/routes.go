package http_handler

import "github.com/gin-gonic/gin"

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/liveness", h.liveness)
	router.GET("/readiness", h.readiness)

	return router
}
