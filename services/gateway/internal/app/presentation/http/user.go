package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Produce json
// @Router /users/ping [get]
func (h *Handler) pingUser(c *gin.Context) {
	message, err := h.user.Ping(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": message,
	})
}

// @Produce json
// @Param id path string true "User ID"
// @Router /users/{id} [get]
func (h *Handler) getUserById(c *gin.Context) {
	userId := c.Param("id")
	result, err := h.user.GetUserById(c.Request.Context(), userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, result)
}
