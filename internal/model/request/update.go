package request

type UpdateRequest struct {
	UserID        uint          `json:"user_id" validate:"required,gte=1"`
	WalletAddress string        `json:"wallet_address,omitempty" validate:"omitempty,min=26,max=66,printascii"`
	Notification  *Notification `json:"notification,omitempty"`
}
