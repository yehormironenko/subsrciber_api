package service

import (
	"context"

	"subsctiption-service/internal/model/request"
	"subsctiption-service/internal/model/response"
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
}
