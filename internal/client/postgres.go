package client

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"

	"subsctiption-service/internal/config"

	"go.uber.org/zap"
)

func CreatePostgresClient(c *config.Config, logger *zap.Logger) *sql.DB {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		c.Postgres.Username, c.Postgres.Password, c.Postgres.Host, c.Postgres.Port, c.Postgres.Database)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		logger.Fatal("Failed to open PostgreSQL connection", zap.Error(err))
	}

	if err := db.Ping(); err != nil {
		logger.Fatal("Failed to connect to PostgreSQL", zap.Error(err))
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(5 * time.Minute)

	logger.Info("Successfully connected to PostgreSQL")
	return db
}
