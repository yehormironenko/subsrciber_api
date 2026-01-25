package response

type Subscriptions struct {
	UserID  uint     `json:"user_id"`
	Wallets []Wallet `json:"wallets"`
}

type Wallet struct {
	Address      string      `json:"address"`
	Preferencies Preferences `json:"preferencies"`
}

type Preferences struct {
	EmailNotifications     *bool `json:"email_notifications,omitempty"`
	WebSocketNotifications *bool `json:"websocket_notifications,omitempty"`
}
