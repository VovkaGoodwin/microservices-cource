package http_server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HttpServer struct {
	server *http.Server
}

func NewHttpServer(handler *gin.Engine) *HttpServer {
	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: handler,
	}

	return &HttpServer{server: server}
}

func (h *HttpServer) Start() error {
	err := h.server.ListenAndServe()
	return err
}
