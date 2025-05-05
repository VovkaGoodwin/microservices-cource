package http

import (
	"gateway/internal/app/application/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

type sendMessageRequest struct {
	UserId string `json:"user_id" binding:"required"`
	Text   string `json:"text" binding:"required"`
}

// @Produce json
// @Router /messages [post]
// @Param message body sendMessageRequest true "сообщение"
func (h *Handler) sendMessage(c *gin.Context) {
	var req sendMessageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sendMessageDto := dto.SendMessageRequest{
		UserId: req.UserId,
		Text:   req.Text,
	}

	result, err := h.message.SendMessage(c.Request.Context(), sendMessageDto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

// @Produce json
// @Router /messages [get]
// @Param user_id query string true "user id"
func (h *Handler) getMessagesByUserId(c *gin.Context) {
	userId := c.Query("user_id")
	if userId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id required"})
		return
	}

	result, err := h.message.GetMessagesByUserId(c.Request.Context(), userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

// @Produce json
// @Router /messages/ping [get]
func (h *Handler) pingMessage(c *gin.Context) {
	message, err := h.message.Ping(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": message})
}
