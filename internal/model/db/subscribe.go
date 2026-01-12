package db

import "time"

type Subscribe struct {
	SubscriptionID int
	UserID         int
	WalletAddress  string
	CreatedAt      time.Time
	Notification   *Notification
}

type Notification struct {
	Email     bool
	WebSocket bool
}
