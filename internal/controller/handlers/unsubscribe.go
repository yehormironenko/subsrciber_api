package handlers

import (
	"net/http"

	"subscription-service/internal/model/request"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (h *Handler) Unsubscriber(c *gin.Context) {
	var unsubscribeRequest request.UnsubscribeRequest

	h.logger.Info("Unsubscribe handler called")

	if err := c.ShouldBindJSON(&unsubscribeRequest); err != nil {
		h.logger.Error("Failed to bind json body", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.validate.Struct(unsubscribeRequest); err != nil {
		h.logger.Error("Failed to validate json body", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	msg, err := h.services.Unsubscribe(c.Request.Context(), unsubscribeRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, msg)
}
