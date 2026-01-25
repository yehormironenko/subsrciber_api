package repository

import (
	"context"

	"subscription-service/internal/model/db"
	"subscription-service/internal/model/request"
)

type RepositoryInterface interface {
	User(ctx context.Context, subscriber request.User) (db.User, error)
	Subscribe(ctx context.Context, subscriber request.SubscribeRequest) (db.Subscribe, error)
	Unsubscribe(ctx context.Context, subscriber request.UnsubscribeRequest) (db.Subscribe, error)
	GetSubscriptions(ctx context.Context, subscriptions request.Subscriptions) (db.Subscriptions, error)
	UpdateSubscriber(ctx context.Context, update request.UpdateRequest) (db.Subscriptions, error)
}
