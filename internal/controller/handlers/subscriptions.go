package handlers

import (
	"net/http"

	"subscription-service/internal/model/request"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (h *Handler) Subscriptions(c *gin.Context) {
	var subscriptions request.Subscriptions

	h.logger.Info("Subscriptions request")

	if err := c.ShouldBindJSON(&subscriptions); err != nil {
		h.logger.Error("Failed to bind json body", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.validate.Struct(subscriptions); err != nil {
		h.logger.Error("Failed to validate json body", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	msg, err := h.services.Subscriptions(c.Request.Context(), subscriptions)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, msg)
}
