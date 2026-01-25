package service

import (
	"context"

	"subscription-service/internal/model/request"
	"subscription-service/internal/model/response"
)

type ServiceInterface interface {
	EchoServiceInterface
	UserServiceInterface
	SubscriberServiceInterface
	SubscriptionsServiceInterface
}

type EchoServiceInterface interface {
	Echo(ctx context.Context) (string, error)
}

type UserServiceInterface interface {
	User(ctx context.Context, user request.User) (response.User, error)
}

type SubscriberServiceInterface interface {
	Subscribe(ctx context.Context, subscriber request.SubscribeRequest) (response.SubscribeResponse, error)
	Unsubscribe(ctx context.Context, unsubscribe request.UnsubscribeRequest) (response.SubscribeResponse, error)
}

type SubscriptionsServiceInterface interface {
	Subscriptions(ctx context.Context, subscriber request.Subscriptions) (response.Subscriptions, error)
	UpdatePreferences(ctx context.Context, update request.UpdateRequest) (response.Subscriptions, error)
}
