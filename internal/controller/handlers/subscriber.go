package handlers

import (
	"net/http"

	"subsctiption-service/internal/model/request"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (h *Handler) Subscriber(c *gin.Context) {
	var subscriber request.SubscribeRequest

	h.logger.Info("Subscribe handler called")

	if err := c.ShouldBindJSON(&subscriber); err != nil {
		h.logger.Error("Failed to bind json body", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.validate.Struct(subscriber); err != nil {
		h.logger.Error("Failed to validate json body", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	msg, err := h.services.Subscribe(c.Request.Context(), subscriber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, msg)
}
