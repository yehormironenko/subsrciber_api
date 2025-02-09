package handlers

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"subsctiption-service/internal/path"
	"subsctiption-service/internal/service"
)

type Handler struct {
	logger   *zap.Logger
	services service.ServiceInterface
}

func NewHandler(services service.ServiceInterface, logger *zap.Logger) *Handler {
	return &Handler{
		services: services,
		logger:   logger,
	}
}

func (h *Handler) Register(r *gin.Engine) {
	r.GET(path.EchoRoute, h.Echo)
	r.POST(path.SubscribeRoute, h.Subscriber)
}
