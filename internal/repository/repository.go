package repository

import (
	"context"
	"database/sql"
	"errors"

	"subscription-service/internal/model/db"
	"subscription-service/internal/model/request"

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
	defer func() {
		if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				r.logger.Error("failed to rollback", zap.Error(rbErr))
			}
		}
	}()

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
		updateQuery := "UPDATE notification_preferences SET email_notifications = $1 WHERE user_id = $2 AND wallet_address = $3 RETURNING email_notifications"
		err = tx.QueryRowContext(ctx, updateQuery, subscriber.Notification.Email, subscriber.UserID, subscriber.WalletAddress).Scan(&dbSubscriber.Notification.Email)
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

func (r *Repository) Unsubscribe(ctx context.Context, unsubscribeRequest request.UnsubscribeRequest) (db.Subscribe, error) {
	r.logger.Info("unsubscribing user", zap.Any("unsubscribe", unsubscribeRequest))
	tx, err := r.dbConn.BeginTx(ctx, nil)
	defer func() {
		if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				r.logger.Error("failed to rollback", zap.Error(rbErr))
			}
		}
	}()

	if err != nil {
		r.logger.Error("failed to start transaction", zap.Error(err))
		return db.Subscribe{}, err
	}

	var dbSubscriber db.Subscribe
	checkQuery := "SELECT id, user_id, wallet_address, created_at FROM wallet_subscriptions WHERE user_id = $1 and wallet_address= $2"
	err = tx.QueryRowContext(ctx, checkQuery, unsubscribeRequest.UserID, unsubscribeRequest.WalletAddress).Scan(&dbSubscriber.SubscriptionID, &dbSubscriber.UserID, &dbSubscriber.WalletAddress, &dbSubscriber.CreatedAt)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		r.logger.Error("failed to check user", zap.Error(err))
		return db.Subscribe{}, err
	}
	// TODO add websocket logic
	if unsubscribeRequest.Notification != nil {
		dbSubscriber.Notification = &db.Notification{}
		updateQuery := "UPDATE notification_preferences SET email_notifications = $1 WHERE user_id = $2 AND wallet_address = $3 RETURNING email_notifications"
		err = tx.QueryRowContext(ctx, updateQuery, unsubscribeRequest.Notification.Email, unsubscribeRequest.UserID, unsubscribeRequest.WalletAddress).Scan(&dbSubscriber.Notification.Email)
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

func (r *Repository) GetSubscriptions(ctx context.Context, subscriptions request.Subscriptions) (db.Subscriptions, error) {
	r.logger.Info("get subscriptions", zap.Any("subscriptions", subscriptions))

	return r.getSubscriptions(ctx, subscriptions)
}

func (r *Repository) getSubscriptions(ctx context.Context, subscriptions request.Subscriptions) (db.Subscriptions, error) {
	r.logger.Info("get subscriptions for user", zap.Any("subscriptions", subscriptions))

	if subscriptions.WalletAddress != "" {
		return r.getSubscriptionForWallet(ctx, subscriptions)
	}

	subsQuery := `SELECT ws.wallet_address,
		COALESCE(np.email_notifications, FALSE) as email_notifications,
		COALESCE(np.websocket_notifications, FALSE) as websocket_notifications
	FROM wallet_subscriptions ws
	LEFT JOIN notification_preferences np
		ON ws.user_id = np.user_id AND ws.wallet_address = np.wallet_address
	WHERE ws.user_id = $1`
	rows, err := r.dbConn.QueryContext(ctx, subsQuery, subscriptions.UserId)
	if err != nil {
		return db.Subscriptions{}, err
	}
	defer rows.Close()

	var result db.Subscriptions

	result.UserID = subscriptions.UserId

	for rows.Next() {
		var wallet db.Wallet

		err = rows.Scan(
			&wallet.Address,
			&wallet.Preferences.Email,
			&wallet.Preferences.Websocket,
		)
		if err != nil {
			return db.Subscriptions{}, err
		}

		result.Wallets = append(result.Wallets, wallet)
	}

	if err = rows.Err(); err != nil {
		return db.Subscriptions{}, err
	}

	return result, nil
}

func (r *Repository) getSubscriptionForWallet(ctx context.Context, subscriptions request.Subscriptions) (db.Subscriptions, error) {
	r.logger.Info("get subscriptions for wallet", zap.Any("subscriptions", subscriptions))

	var dbSubscriber db.Subscriptions
	var wallet db.Wallet

	subsQuery := `SELECT ws.user_id, ws.wallet_address,
		COALESCE(np.email_notifications, FALSE) as email_notifications,
		COALESCE(np.websocket_notifications, FALSE) as websocket_notifications
	FROM wallet_subscriptions ws
	LEFT JOIN notification_preferences np
		ON ws.user_id = np.user_id AND ws.wallet_address = np.wallet_address
	WHERE ws.user_id = $1 AND ws.wallet_address = $2`

	row := r.dbConn.QueryRowContext(ctx, subsQuery, subscriptions.UserId, subscriptions.WalletAddress)
	err := row.Scan(&dbSubscriber.UserID, &wallet.Address, &wallet.Preferences.Email, &wallet.Preferences.Websocket)
	if errors.Is(err, sql.ErrNoRows) {
		return db.Subscriptions{UserID: subscriptions.UserId, Wallets: []db.Wallet{}}, nil
	}
	if err != nil {
		r.logger.Error("failed to get user", zap.Error(err))
		return db.Subscriptions{}, err
	}
	dbSubscriber.Wallets = append(dbSubscriber.Wallets, wallet)
	return dbSubscriber, nil
}

func (r *Repository) UpdateSubscriber(ctx context.Context, update request.UpdateRequest) (db.Subscriptions, error) {
	r.logger.Info("update subscriptions", zap.Any("update", update))

	if update.Notification == nil {
		return db.Subscriptions{}, errors.New("notification preferences required")
	}

	tx, err := r.dbConn.BeginTx(ctx, nil)
	if err != nil {
		r.logger.Error("failed to start transaction", zap.Error(err))
		return db.Subscriptions{}, err
	}
	defer tx.Rollback()

	updateQuery := `UPDATE notification_preferences
                   SET email_notifications = $1, websocket_notifications = $2
                   WHERE user_id = $3`
	args := []interface{}{
		update.Notification.Email,
		update.Notification.Websocket,
		update.UserID,
	}

	if update.WalletAddress != "" {
		updateQuery += " AND wallet_address = $4"
		args = append(args, update.WalletAddress)
	}

	result, err := tx.ExecContext(ctx, updateQuery, args...)
	if err != nil {
		r.logger.Error("failed to update notification preferences", zap.Error(err))
		return db.Subscriptions{}, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		r.logger.Error("failed to get rows affected", zap.Error(err))
		return db.Subscriptions{}, err
	}

	if rowsAffected == 0 {
		r.logger.Warn("no records updated", zap.Uint("user_id", update.UserID))
		return db.Subscriptions{}, errors.New("no records updated")
	}

	err = tx.Commit()
	if err != nil {
		r.logger.Error("failed to commit transaction", zap.Error(err))
		return db.Subscriptions{}, err
	}

	// Вернуть обновленные данные
	subscriptions := request.Subscriptions{
		UserId:        update.UserID,
		WalletAddress: update.WalletAddress,
	}
	return r.getSubscriptions(ctx, subscriptions)
}
