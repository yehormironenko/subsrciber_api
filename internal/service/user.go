package service

import (
	"context"
	"time"

	"go.uber.org/zap"
	"subscription-service/internal/model/db"
	"subscription-service/internal/model/request"
	"subscription-service/internal/model/response"
	"subscription-service/internal/repository"
)

type User struct {
	repo   repository.RepositoryInterface
	logger *zap.Logger
}

func NewUserService(repository repository.RepositoryInterface, logger *zap.Logger) *User {
	return &User{
		repo:   repository,
		logger: logger,
	}
}

func (s *User) User(ctx context.Context, user request.User) (response.User, error) {
	s.logger.Info("user service called")
	dbUser, err := s.repo.User(ctx, user)
	if err != nil {
		s.logger.Error("user service failed", zap.Error(err))
		return response.User{}, err
	}

	return dbUserModelToJsonModel(dbUser), nil
}

func dbUserModelToJsonModel(subscriber db.User) response.User {
	return response.User{
		UserId:    subscriber.UserID,
		Email:     subscriber.Email,
		CreatedAt: subscriber.CreateTime.Format(time.RFC3339),
	}
}
