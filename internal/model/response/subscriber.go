package response

type SubscribeResponse struct {
	SubscriptionId int          `json:"subscription_id"`
	UserId         int          `json:"user_id"`
	WalletAddress  string       `json:"wallet_address"`
	CreatedAt      string       `json:"created_at"`
	Notification   Notification `json:"notification"`
}

type Notification struct {
	Email     *bool `json:"email,omitempty"`
	WebSocket *bool `json:"web_socket,omitempty"`
}
