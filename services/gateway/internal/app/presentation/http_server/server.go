package http_server

import "net/http"

type HttpServer struct {
	s *http.Server
}

func NewHttpServer(handler *Handler) *HttpServer {
	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: handler.InitRoutes(),
	}

	return &HttpServer{
		s: server,
	}
}

func (h *HttpServer) Start() error {
	err := h.s.ListenAndServe()
	return err
}
