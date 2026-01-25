package handlers

import (
	"net/http"

	"subscription-service/internal/model/request"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (h *Handler) User(c *gin.Context) {
	var user request.User

	h.logger.Info("User handler called")

	if err := c.ShouldBindJSON(&user); err != nil {
		h.logger.Error("Failed to bind json body", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.validate.Struct(user); err != nil {
		h.logger.Error("Failed to validate json body", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	msg, err := h.services.User(c.Request.Context(), user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, msg)
}
