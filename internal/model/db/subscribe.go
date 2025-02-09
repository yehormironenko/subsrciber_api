package db

import "time"

type Subscribe struct {
	UserID     int
	Email      string
	CreateTime time.Time
}
