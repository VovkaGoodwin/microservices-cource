package http_server

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.GET("/check_token", h.checkToken)
		auth.POST("/authenticate", h.authenticate)
		auth.GET("/ping", h.pingAuth)
	}

	user := router.Group("/users")
	{
		user.GET("/:id", h.getUserById)
		user.GET("/ping", h.pingUser)
	}

	message := router.Group("/messages")
	{
		message.GET("", h.getMessagesByUserId)
		message.GET("/ping", h.pingMessage)
		message.POST("", h.sendMessage)
	}

	swagger := router.Group("/swagger")
	{
		swagger.GET("/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	return router
}
