package db

import "time"

type User struct {
	UserID     int
	Email      string
	CreateTime time.Time
}
