package http_server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Produce json
// @Router /auth/check_token [get]
// @Param token query string true "токен на проверку"
func (h *Handler) checkToken(c *gin.Context) {
	valid, userId, err := h.auth.CheckToken(c.Request.Context(), c.Query("token"))
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

type authRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) authenticate(c *gin.Context) {
	var body authRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	token, err := h.auth.Authenticate(c.Request.Context(), body.Username, body.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
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
