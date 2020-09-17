package account

import (
	"condomilux/condo-admin/internal/interfaces"
)

// Opts in a comment
type Opts struct {
	Repository interfaces.Account
}

// CreateRequest in a comment
type CreateRequest struct {
	ID            string `json:"account_id" sql:"account_id"`
	BankName      string `json:"bank" sql:"bank"`
	AccountNumber string `json:"number" sql:"number"`
	Holder        string `json:"holder" sql:"holder"`
}

// UpdateRequest in a comment
type UpdateRequest CreateRequest
