package request

type SubscribeRequest struct {
	UserID        uint          `json:"userID" validate:"required,gte=1"`
	WalletAddress string        `json:"walletAddress" validate:"required,min=26,max=66,printascii"`
	Notification  *Notification `json:"notification,omitempty"`
}

type Notification struct {
	Email     bool `json:"email,omitempty"`
	WebSocket bool `json:"webSocket,omitempty"`
}
