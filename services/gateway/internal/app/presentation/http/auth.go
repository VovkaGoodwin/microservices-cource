package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Produce json
// @Router /auth/check_token [get]
// @Param token query string true "токен на проверку"
func (h *Handler) checkToken(c *gin.Context) {
	valid, userId, err := h.auth.CheckToken(c.Request.Context(), c.Param("token"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	var message = "false"
	if valid {
		message = "ok"
	}

	c.JSON(http.StatusOK, gin.H{
		"valid":  message,
		"userId": userId,
	})
}

// @Produce json
// @Router /auth/ping [get]
func (h *Handler) pingAuth(c *gin.Context) {
	message, err := h.auth.Ping(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": message,
	})
}
