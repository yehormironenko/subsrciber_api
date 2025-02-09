package handlers

import (
	"github.com/gin-gonic/gin"
)

func (h *Handler) Echo(c *gin.Context) {

	h.logger.Info("Echo handler called")

	msg, err := h.services.Echo(c.Request.Context())
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, msg)
}
