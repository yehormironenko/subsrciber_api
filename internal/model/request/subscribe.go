package request

type SubscribeRequest struct {
	UserID        uint   `validate:"required,gte=1"`
	WalletAddress string `validate:"required,gte=1,lte=50"`
}
