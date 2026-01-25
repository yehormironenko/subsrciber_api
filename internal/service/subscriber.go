package service

import (
	"context"
	"time"

	"subscription-service/internal/model/db"
	"subscription-service/internal/model/request"
	"subscription-service/internal/model/response"
	"subscription-service/internal/repository"

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

func (s *Subscriber) Unsubscribe(ctx context.Context, unsubscribe request.UnsubscribeRequest) (response.SubscribeResponse, error) {
	s.logger.Info("unsubscriber service called")
	sub, err := s.repository.Unsubscribe(ctx, unsubscribe)
	if err != nil {
		s.logger.Error("unsubscriber service failed", zap.Error(err))
		return response.SubscribeResponse{}, err
	}

	return dbSubscriberModelToJsonModel(sub), nil
}

func dbSubscriberModelToJsonModel(subscriber db.Subscribe) response.SubscribeResponse {
	resp := response.SubscribeResponse{
		SubscriptionId: subscriber.SubscriptionID,
		UserId:         subscriber.UserID,
		WalletAddress:  subscriber.WalletAddress,
		CreatedAt:      subscriber.CreatedAt.Format(time.RFC3339),
	}

	if subscriber.Notification != nil {
		resp.Notification = response.Notification{
			Email:     &subscriber.Notification.Email,
			WebSocket: &subscriber.Notification.WebSocket,
		}
	}

	return resp
}
