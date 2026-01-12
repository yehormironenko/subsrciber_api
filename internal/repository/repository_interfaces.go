package repository

import (
	"context"

	"subsctiption-service/internal/model/db"
	"subsctiption-service/internal/model/request"
)

type RepositoryInterface interface {
	User(ctx context.Context, subscriber request.User) (db.User, error)
	Subscribe(ctx context.Context, subscriber request.SubscribeRequest) (db.Subscribe, error)
	//Unsubscribe
	//UpdateSubscriber
	//GetSubscriptions
}
