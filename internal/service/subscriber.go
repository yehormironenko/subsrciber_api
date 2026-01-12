package service

import (
	"context"
	"time"

	"subsctiption-service/internal/model/db"
	"subsctiption-service/internal/model/request"
	"subsctiption-service/internal/model/response"
	"subsctiption-service/internal/repository"

	"go.uber.org/zap"
)

type Subscriber struct {
	repository repository.RepositoryInterface
	logger     *zap.Logger
}

func NewSubscriber(repository repository.RepositoryInterface, logger *zap.Logger) *Subscriber {
	return &Subscriber{
		repository: repository,
		logger:     logger,
	}
}

func (s *Subscriber) Subscribe(ctx context.Context, subscriber request.SubscribeRequest) (response.SubscribeResponse, error) {
	s.logger.Info("subscriber service called")
	sub, err := s.repository.Subscribe(ctx, subscriber)
	if err != nil {
		s.logger.Error("subscriber service failed", zap.Error(err))
		return response.SubscribeResponse{}, err
	}

	return dbSubscriberModelToJsonModel(sub), nil
}

func dbSubscriberModelToJsonModel(subscriber db.Subscribe) response.SubscribeResponse {
	return response.SubscribeResponse{
		SubscriptionId: subscriber.SubscriptionID,
		UserId:         subscriber.UserID,
		WalletAddress:  subscriber.WalletAddress,
		CreatedAt:      subscriber.CreatedAt.Format(time.RFC3339),
		Notification: response.Notification{
			Email:     &subscriber.Notification.Email,
			WebSocket: &subscriber.Notification.WebSocket,
		},
	}
}
