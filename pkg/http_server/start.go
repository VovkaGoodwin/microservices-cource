package http_server

import (
	"context"
	"net"
)

func (h *httpServer) Start(ctx context.Context) error {
	h.server.ConnContext = func(_ context.Context, _ net.Conn) context.Context {
		return ctx
	}
	err := h.server.ListenAndServe()
	return err
}
