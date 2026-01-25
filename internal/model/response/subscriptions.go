package response

type Subscriptions struct {
	UserID  uint     `json:"user_id"`
	Wallets []Wallet `json:"wallets"`
}

type Wallet struct {
	Address      string       `json:"address"`
	Preferencies Preferencies `json:"preferencies"`
}

type Preferencies struct {
	EmailNotifications     *bool `json:"email_notifications,omitempty"`
	WebSocketNotifications *bool `json:"websocket_notifications,omitempty"`
}
