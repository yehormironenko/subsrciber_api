package service

import (
	"context"

	"subsctiption-service/internal/model/request"
	"subsctiption-service/internal/model/response"
)

type ServiceInterface interface {
	EchoServiceInterface
	SubscriberServiceInterface
}

type EchoServiceInterface interface {
	Echo(ctx context.Context) (string, error)
}

type SubscriberServiceInterface interface {
	Subscribe(ctx context.Context, user request.Subscriber) (response.Subscriber, error)
}
