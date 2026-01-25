package db

type Subscriptions struct {
	UserID  uint
	Wallets []Wallet
}

type Wallet struct {
	Address     string
	Preferences Preferences
}

type Preferences struct {
	Email     bool
	Websocket bool
}
