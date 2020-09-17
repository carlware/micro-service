package models

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Account defines what information the user inputted for his/her payments
type Account struct {
	ID            string    `json:"id" sql:"id"`
	BankName      string    `json:"bank" sql:"bank"`
	AccountNumber string    `json:"number" sql:"number"`
	Holder        string    `json:"holder" sql:"holder"`
	Created       time.Time `json:"created" sql:"created"`
}

// NewAccount creates a new instance of Account
func NewAccount(bank, number, holder string) *Account {
	return &Account{
		ID:            "",
		BankName:      bank,
		AccountNumber: number,
		Holder:        holder,
	}
}

// Initialize generates a new ID for Account
func (a *Account) Initialize() {
	a.ID = IDGeneratorFunc()
	a.Created = time.Now()
}

// Validate tests for some fields to have correct values
func (a Account) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.BankName, validation.Required, validation.Length(1, 255)),
		validation.Field(&a.AccountNumber, validation.Required, validation.Length(18, 18)),
		validation.Field(&a.Holder, validation.Required, validation.Length(1, 255)),
	)
}

// Id is...
func (a *Account) Id() string {
	return a.ID
}
