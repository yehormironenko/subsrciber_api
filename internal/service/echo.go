package service

import (
	"context"

	"go.uber.org/zap"
)

type Echo struct {
	logger *zap.Logger
}

func NewEchoService(logger *zap.Logger) *Echo {
	return &Echo{
		logger: logger,
	}
}

func (s *Echo) Echo(_ context.Context) (string, error) {
	s.logger.Info("Echo service called")
	return "Echo service", nil
}
