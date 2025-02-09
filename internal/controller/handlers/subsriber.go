package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"subsctiption-service/internal/model/request"
)

func (h *Handler) Subscriber(c *gin.Context) {
	var user request.Subscriber

	h.logger.Info("Subscriber handler called")

	if err := c.ShouldBindJSON(&user); err != nil {
		h.logger.Error("Failed to bind json body", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		h.logger.Error("Failed to validate json body", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	msg, err := h.services.Subscribe(c.Request.Context(), user)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, msg)
}
