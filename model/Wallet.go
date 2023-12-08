package model

type Wallet struct {
	UserCode  string `json:"user_code"`
	Balance   int64  `json:"balance"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
