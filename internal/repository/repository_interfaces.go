package repository

import (
	"context"

	"subsctiption-service/internal/model/db"
	"subsctiption-service/internal/model/request"
)

type RepositoryInterface interface {
	Subscribe(ctx context.Context, subscriber request.Subscriber) (db.Subscribe, error)
	//Unsubscribe
	//UpdateSubscriber
	//GetSubscriptions
}
