package service

import (
	"context"

	"subsctiption-service/internal/model/request"
	"subsctiption-service/internal/model/response"
)

type ServiceInterface interface {
	EchoServiceInterface
	UserServiceInterface
}

type EchoServiceInterface interface {
	Echo(ctx context.Context) (string, error)
}

type UserServiceInterface interface {
	User(ctx context.Context, user request.User) (response.User, error)
}
