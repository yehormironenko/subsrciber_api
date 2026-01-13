package handlers

import (
	"subsctiption-service/internal/path"
	"subsctiption-service/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

type Handler struct {
	logger   *zap.Logger
	services service.ServiceInterface
	validate *validator.Validate
}

func NewHandler(services service.ServiceInterface, logger *zap.Logger) *Handler {
	return &Handler{
		services: services,
		logger:   logger,
		validate: validator.New(),
	}
}

func (h *Handler) Register(r *gin.Engine) {
	r.GET(path.EchoRoute, h.Echo)
	r.POST(path.CreateUserRoute, h.User)
	r.POST(path.SubscribeRoute, h.Subscriber)
	r.POST(path.UnsubscribeRoute, h.Unsubscriber)
}
