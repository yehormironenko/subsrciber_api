package request

type UnsubscribeRequest struct {
	UserID        uint          `json:"user_id" validate:"required,gte=1"`
	WalletAddress string        `json:"wallet_address" validate:"required,min=26,max=66,printascii"`
	Notification  *Notification `json:"notification,omitempty"`
}
