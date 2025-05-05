package http_server

import (
	"github.com/VovkaGoodwin/microservices-cource/pkg/app"
	"github.com/VovkaGoodwin/microservices-cource/pkg/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HttpServer interface {
	app.Runner
}

type httpServer struct {
	server *http.Server
}

var _ HttpServer = (*httpServer)(nil)

func NewHttpServer(cfg *config.Config, handler *gin.Engine) HttpServer {
	server := &http.Server{
		Addr:         cfg.HTTPServer.Address,
		ReadTimeout:  cfg.HTTPServer.Timeout,
		WriteTimeout: cfg.HTTPServer.Timeout,
		IdleTimeout:  cfg.HTTPServer.IdleTimeout,
		Handler:      handler,
	}

	return &httpServer{server: server}
}
