package http_handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) liveness(c *gin.Context) {
	r := h.hc.LivenessProbe(c.Request.Context())
	h.respondSuccess(c, http.StatusOK, r.Result)
}

func (h *Handler) readiness(c *gin.Context) {
	r := h.hc.ReadinessProbe(c.Request.Context())
	h.respondSuccess(c, http.StatusOK, r.Result)
}
