package db

import "time"

type User struct {
	UserID     uint
	Email      string
	CreateTime time.Time
}
