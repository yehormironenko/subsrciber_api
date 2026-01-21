package request

type Subscriptions struct {
	UserId        string `json:"user_id"`
	WalletAddress string `json:"wallet_address,omitzero"`
}
