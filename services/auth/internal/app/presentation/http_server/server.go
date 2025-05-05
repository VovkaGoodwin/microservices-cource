package http_server

import (
	"net/http"

	"github.com/VovkaGoodwin/microservices-cource/pkg/config"
	"github.com/gin-gonic/gin"
)

type HttpServer struct {
	server *http.Server
}

func NewHttpServer(cfg *config.Config, handler *gin.Engine) *HttpServer {
	server := &http.Server{
		Addr:         cfg.HTTPServer.Address,
		ReadTimeout:  cfg.HTTPServer.Timeout,
		WriteTimeout: cfg.HTTPServer.Timeout,
		IdleTimeout:  cfg.HTTPServer.IdleTimeout,
		Handler:      handler,
	}

	return &HttpServer{server: server}
}

func (h *HttpServer) Start() error {
	err := h.server.ListenAndServe()
	return err
}
