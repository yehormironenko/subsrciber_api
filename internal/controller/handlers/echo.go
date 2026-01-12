package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Echo(c *gin.Context) {

	h.logger.Info("Echo handler called")

	msg, err := h.services.Echo(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, msg)
}
