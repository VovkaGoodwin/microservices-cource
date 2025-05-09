package http_server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) liveness(c *gin.Context) {
	r := h.hc.LivenessProbe(c.Request.Context())
	h.respondSuccess(c, http.StatusOK, r.Result)
}

func (h *Handler) readiness(c *gin.Context) {
	r := h.hc.ReadinessProbe(c.Request.Context())
	if !r.Success {
		h.respondError(c, http.StatusServiceUnavailable, r.Result)
		return
	}

	h.respondSuccess(c, http.StatusOK, r.Result)
}
