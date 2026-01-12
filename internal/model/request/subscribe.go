package request

type SubscribeRequest struct {
	UserID        uint          `json:"userID" validate:"required,gte=1"`
	WalletAddress string        `json:"walletAddress" validate:"required,gte=1,lte=50"`
	Notification  *Notification `json:"notification,omitempty"`
}

type Notification struct {
	Email     bool `json:"email,omitempty"`
	WebSocket bool `json:"webSocket,omitempty"` // TODO - implement webSocket notification
}
