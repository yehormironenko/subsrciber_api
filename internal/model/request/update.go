package request

type UpdateRequest struct {
	UserID        uint          `json:"userID" validate:"required,gte=1"`
	WalletAddress string        `json:"walletAddress,omitempty" validate:"omitempty,min=26,max=66,printascii"`
	Notification  *Notification `json:"notification,omitempty"`
}
