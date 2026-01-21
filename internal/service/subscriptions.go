package service

import (
	"context"

	"subsctiption-service/internal/model/db"
	"subsctiption-service/internal/model/request"
	"subsctiption-service/internal/model/response"
	"subsctiption-service/internal/repository"

	"go.uber.org/zap"
)

type Subscriptions struct {
	repository repository.RepositoryInterface
	logger     *zap.Logger
}

func NewSubscriptions(repository repository.RepositoryInterface, logger *zap.Logger) *Subscriptions {
	return &Subscriptions{
		repository: repository,
		logger:     logger,
	}
}

func (s *Subscriptions) Subscriptions(ctx context.Context, subscriptions request.Subscriptions) (response.Subscriptions, error) {
	s.logger.Info("get subscriptions service called")
	sub, err := s.repository.GetSubscriptions(ctx, subscriptions)
	if err != nil {
		s.logger.Error("subscriber service failed", zap.Error(err))
		return response.Subscriptions{}, err
	}

	return dbSubscriptionsToJsonModel(sub), nil
}

func dbSubscriptionsToJsonModel(subs db.Subscriptions) response.Subscriptions {
	var wallets []response.Wallet
	for _, v := range subs.Wallets {
		wallets = append(wallets, response.Wallet{Address: v.Address,
			Preferencies: response.Preferencies{EmailNotifications: &v.Preferences.Email, WebSocketNotifications: &v.Preferences.Websocket}})
	}

	return response.Subscriptions{
		UserID:  subs.UserID,
		Wallets: wallets,
	}
}
