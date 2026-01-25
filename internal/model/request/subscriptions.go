package request

type Subscriptions struct {
	UserId        uint   `json:"user_id"`
	WalletAddress string `json:"wallet_address,omitzero"`
}
