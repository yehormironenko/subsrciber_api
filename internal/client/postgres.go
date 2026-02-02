package client

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"time"

	_ "github.com/lib/pq"

	"subscription-service/internal/config"

	"go.uber.org/zap"
)

func CreatePostgresClient(c *config.Config, logger *zap.Logger) *sql.DB {
	var (
		username = os.Getenv("POSTGRES_USERNAME")
		password = os.Getenv("POSTGRES_PASSWORD")
		host     = os.Getenv("POSTGRES_HOST")
		port     = os.Getenv("POSTGRES_PORT")
		database = os.Getenv("POSTGRES_NAME")
	)

	if username == "" {
		username = c.Postgres.Username
	}

	if password == "" {
		password = c.Postgres.Password
	}

	if host == "" {
		host = c.Postgres.Host
	}

	if port == "" {
		port = strconv.Itoa(c.Postgres.Port)
	}

	if database == "" {
		database = c.Postgres.Database
	}

	sslmode := os.Getenv("POSTGRES_SSL_MODE")
	if sslmode == "" {
		sslmode = "disable" // default для dev
	}

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		username, password, host, port, database, sslmode)

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
