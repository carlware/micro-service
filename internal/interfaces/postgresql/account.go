package postgresql

import (
	"condomilux/condo-admin/internal/helpers"
	"condomilux/condo-admin/internal/interfaces"
	"condomilux/condo-admin/internal/models"
	"context"

	"github.com/go-pg/pg"
)

type account struct {
	db *pg.DB
}

func NewAccount(db *pg.DB) interfaces.Account {
	return &account{
		db: db,
	}
}

// TODO: analyze to reduce the output arguments
func (p *account) Add(ctx context.Context, account *models.Account) (*models.Account, error) {
	// TODO: Invalidate cache
	const op = "interfaces.account.Insert"
	err := p.db.Insert(account)
	if err != nil {
		return nil, helpers.NewDatabaseError(op, err)
	}
	return account, nil
}

func (p *account) Remove(ctx context.Context, account *models.Account) (*models.Account, error) {
	// TODO: Invalidate cache
	const op = "interfaces.account.Remove"
	err := p.db.Delete(account)
	if err != nil {
		return nil, helpers.NewDatabaseError(op, err)
	}
	return account, nil
}

func (p *account) Get(ctx context.Context, accountID string) (*models.Account, error) {
	// TODO: set cache key
	const op = "interfaces.account.Get"
	var obj models.Account

	// TODO:  sanitize accountID
	_, err := p.db.QueryOne(&obj, `SELECT * FROM accounts WHERE id = ?`, accountID)
	if err != nil {
		return nil, helpers.NewDatabaseError(op, err)
	}

	return &obj, nil
}

func (p *account) Update(ctx context.Context, account *models.Account) (*models.Account, error) {
	// TODO: set cache key
	const op = "interfaces.account.Update"
	err := p.db.Update(account)
	if err != nil {
		return nil, helpers.NewDatabaseError(op, err)
	}
	return account, nil
}

func (p *account) List(ctx context.Context) ([]*models.Account, error) {
	// TODO: set cache key
	const op = "interfaces.account.List"
	var objs []*models.Account

	_, err := p.db.Query(&objs, `SELECT * FROM accounts`)
	if err != nil {
		return nil, helpers.NewDatabaseError(op, err)
	}

	return objs, nil
}
