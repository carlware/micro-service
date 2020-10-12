package memorydb

import (
	"carlware/accounts/internal/interfaces"
	"carlware/accounts/internal/models"
	"context"

	"github.com/carlware/go-common/errors"
)

type account struct {
}

func NewAccount() interfaces.Account {
	return &account{
	}
}

var DB = map[string]*models.Account{}


func (p *account) Add(ctx context.Context, account *models.Account) (*models.Account, error) {
	const op = "interfaces.account.Insert"
	DB[account.ID] = account
	return account, nil
}

func (p *account) Remove(ctx context.Context, account *models.Account) (*models.Account, error) {
	const op = "interfaces.account.Remove"
	if _, ok := DB[account.ID]; ok {
		delete(DB, account.ID)
		return account, nil	
	} else {
		return nil, errors.New(errors.NotFound, op, "A resource with this ID does not exists", nil)
	}
	
}

func (p *account) Get(ctx context.Context, accountID string) (*models.Account, error) {
	const op = "interfaces.account.Get"
	if val, ok := DB[accountID]; ok {
		return val, nil	
	} else {
		return nil, errors.New(errors.NotFound, op, "A resource with this ID does not exists", nil)
	}
}

func (p *account) Update(ctx context.Context, account *models.Account) (*models.Account, error) {
	const op = "interfaces.account.Update"
	if _, ok := DB[account.ID]; ok {
		DB[account.ID] = account
		return account, nil	
	} else {
		return nil, errors.New(errors.NotFound, op, "A resource with this ID does not exists", nil)
	}
}

func (p *account) List(ctx context.Context) ([]*models.Account, error) {
	accounts := []*models.Account{}
    for _, account := range DB {
        accounts = append(accounts, account)
    }
    return accounts, nil
}
