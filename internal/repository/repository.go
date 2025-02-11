package repository

import (
	"context"
	"database/sql"

	"go.uber.org/zap"
	"subsctiption-service/internal/model/db"
	"subsctiption-service/internal/model/request"
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
	var subscriber db.User

	err := r.dbConn.QueryRowContext(ctx, query, user.Email).Scan(&subscriber.UserID, &subscriber.Email, &subscriber.CreateTime)
	if err != nil {
		r.logger.Error("failed to insert user", zap.Error(err))
		return db.User{}, err
	}

	return subscriber, nil
}
