package service

import (
	"context"
	"time"

	"go.uber.org/zap"
	"subsctiption-service/internal/model/db"
	"subsctiption-service/internal/model/request"
	"subsctiption-service/internal/model/response"
	"subsctiption-service/internal/repository"
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
	sub, err := s.repo.User(ctx, user)
	if err != nil {
		s.logger.Error("user service failed", zap.Error(err))
		return response.User{}, err
	}

	return dbModelToJsonModel(sub), nil
}

func dbModelToJsonModel(subscriber db.User) response.User {
	return response.User{
		UserId:    subscriber.UserID,
		Email:     subscriber.Email,
		CreatedAt: subscriber.CreateTime.Format(time.RFC3339),
	}
}
