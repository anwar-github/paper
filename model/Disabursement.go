package model

import "time"

type Disbursement struct {
	Order         string    `json:"order_id"`
	Success       bool      `json:"success"`
	UserCode      string    `json:"user_code"`
	AccountNumber string    `json:"account_number"`
	Amount        int64     `json:"amount"`
	Balance       int64     `json:"balance"`
	Provider      string    `json:"provider"`
	Request       string    `json:"request_id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
