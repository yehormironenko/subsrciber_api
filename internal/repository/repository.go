package repository

import (
	"context"
	"database/sql"
	"errors"

	"subsctiption-service/internal/model/db"
	"subsctiption-service/internal/model/request"

	"go.uber.org/zap"
)

type Repository struct {
	logger *zap.Logger
	dbConn *sql.DB
}

func NewRepository(db *sql.DB, logger *zap.Logger) RepositoryInterface {
	return &Repository{
		logger: logger,
		dbConn: db,
	}
}
func (r *Repository) User(ctx context.Context, user request.User) (db.User, error) {
	r.logger.Info("creating user", zap.Any("user", user))

	query := "INSERT INTO users (email) VALUES ($1) RETURNING id, email, created_at"
	var dbUser db.User

	err := r.dbConn.QueryRowContext(ctx, query, user.Email).Scan(&dbUser.UserID, &dbUser.Email, &dbUser.CreateTime)
	if err != nil {
		r.logger.Error("failed to insert user", zap.Error(err))
		return db.User{}, err
	}

	return dbUser, nil
}

func (r *Repository) Subscribe(ctx context.Context, subscriber request.SubscribeRequest) (db.Subscribe, error) {
	r.logger.Info("subscribing user", zap.Any("subscriber", subscriber))
	tx, err := r.dbConn.BeginTx(ctx, nil)
	if err != nil {
		r.logger.Error("failed to start transaction", zap.Error(err))
		return db.Subscribe{}, err
	}

	var dbSubscriber db.Subscribe
	checkQuery := "SELECT id, user_id, wallet_address, created_at FROM wallet_subscriptions WHERE user_id = $1 and wallet_address= $2"
	err = tx.QueryRowContext(ctx, checkQuery, subscriber.UserID, subscriber.WalletAddress).Scan(&dbSubscriber.SubscriptionID, &dbSubscriber.UserID, &dbSubscriber.WalletAddress, &dbSubscriber.CreatedAt)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		r.logger.Error("failed to check user", zap.Error(err))
		return db.Subscribe{}, err
	}

	if dbSubscriber.UserID == 0 {
		query := "INSERT INTO wallet_subscriptions (user_id, wallet_address) VALUES ($1,$2) RETURNING id, user_id, wallet_address, created_at"

		err = tx.QueryRowContext(ctx, query, subscriber.UserID, subscriber.WalletAddress).Scan(&dbSubscriber.SubscriptionID, &dbSubscriber.UserID, &dbSubscriber.WalletAddress, &dbSubscriber.CreatedAt)
		if err != nil {
			r.logger.Error("failed to insert user", zap.Error(err))
			return db.Subscribe{}, err
		}
	}
	// TODO add websocket logic
	if subscriber.Notification != nil {
		dbSubscriber.Notification = &db.Notification{}
		updateQuery := "UPDATE notification_preferences SET email_notifications = $1 WHERE user_id = $2 RETURNING email_notifications"
		err = tx.QueryRowContext(ctx, updateQuery, subscriber.Notification.Email, subscriber.UserID).Scan(&dbSubscriber.Notification.Email)
		if err != nil {
			r.logger.Error("failed to update user", zap.Error(err))
			return db.Subscribe{}, err
		}
	}

	err = tx.Commit()
	if err != nil {
		r.logger.Error("failed to commit transaction", zap.Error(err))
		return db.Subscribe{}, err
	}

	return dbSubscriber, nil
}
