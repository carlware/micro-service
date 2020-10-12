// +build unit

package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewAccount(t *testing.T) {
	type args struct {
		bank   string
		number string
		holder string
	}
	tests := []struct {
		name string
		args func(t *testing.T) args

		want1 *Account
	}{
		{
			name: "Testing default",
			args: func(t *testing.T) args {
				return args{
					bank:   "HSBC",
					number: "123456",
					holder: "Person1",
				}
			},
			want1: &Account{
				BankName:      "HSBC",
				AccountNumber: "123456",
				Holder:        "Person1",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tArgs := tt.args(t)

			got1 := NewAccount(tArgs.bank, tArgs.number, tArgs.holder)

			assert.Equal(t, tt.want1, got1)
		})
	}
}

func TestAccount_Validate(t *testing.T) {
	tests := []struct {
		name    string
		model   *Account
		inspect func(r *Account, t *testing.T)

		wantErr    bool
		inspectErr func(err error, t *testing.T)
	}{
		{
			name: "Valid parameters",
			model: &Account{
				BankName:      "BBVA",
				AccountNumber: "263845194802841049",
				Holder:        "Person1",
			},
			wantErr: false,
		},
		{
			name: "Test case where BankName is empty",
			model: &Account{
				BankName:      "",
				AccountNumber: "263845194802841049",
				Holder:        "P1",
			},
			wantErr: true,
			inspectErr: func(err error, t *testing.T) {
				assert.EqualError(t, err, "bank: cannot be blank.")
			},
		},
		{
			name: "test case where AccountNumber is invalid",
			model: &Account{
				BankName:      "BBVA",
				AccountNumber: "1234",
				Holder:        "P2",
			},
			wantErr: true,
			inspectErr: func(err error, t *testing.T) {
				assert.EqualError(t, err, "number: the length must be exactly 18.")
			},
		},
		{
			name: "Test case where Holder field is empty",
			model: &Account{
				BankName:      "BBVA",
				AccountNumber: "27ed37ch1903ks2801",
				Holder:        "",
			},
			wantErr: true,
			inspectErr: func(err error, t *testing.T) {
				assert.EqualError(t, err, "holder: cannot be blank.")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			receiver := tt.model
			err := receiver.Validate()

			if tt.inspect != nil {
				tt.inspect(receiver, t)
			}

			if tt.wantErr {
				require.Error(t, err)
				if tt.inspectErr != nil {
					tt.inspectErr(err, t)
				}
			}
			if !tt.wantErr && err != nil {
				require.Nil(t, err)
			}
		})
	}
}

func TestAccount_Initialize(t *testing.T) {
	tests := []struct {
		name    string
		model   *Account
		inspect func(r *Account, t *testing.T)
	}{
		{
			name: "Test Initialize",
			model: &Account{
				BankName:      "HSBC",
				AccountNumber: "123456",
				Holder:        "Person1",
			},
			inspect: func(r *Account, t *testing.T) {
				assert.NotNil(t, r.ID)
				assert.NotNil(t, r.Created)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			receiver := tt.model
			receiver.Initialize()

			if tt.inspect != nil {
				tt.inspect(receiver, t)
			}

		})
	}
}
