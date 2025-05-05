package app

import (
	"gateway/internal/app/application/interactors"
	"gateway/internal/app/presentation/http"
	"github.com/gin-gonic/gin"
	globalhttp "net/http"
)

func initHttpHandler() *gin.Engine {
	handler := http.NewHandler(
		interactors.NewAuth(GetAuthClient()),
		interactors.NewUser(GetUserClient()),
		interactors.NewMessage(GetMessageClient()),
	)

	return handler.InitRoutes()
}

func initHttpServer(handler globalhttp.Handler) *globalhttp.Server {
	server := &globalhttp.Server{
		Addr:    "0.0.0.0:8080",
		Handler: handler,
	}

	return server
}
