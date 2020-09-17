package dispatchers

import (
	"condomilux/condo-admin/cli/config"
	"condomilux/condo-admin/internal/interfaces"
	"condomilux/condo-admin/internal/interfaces/postgresql"

	"github.com/go-pg/pg"
)

type Repositories struct {
	Account             interfaces.Account
	Accountant          interfaces.Accountant
	Discount            interfaces.Discount
	Deposit             interfaces.Deposit
	Provider            interfaces.Provider
	House               interfaces.House
	Condo               interfaces.Condo
	Expense             interfaces.Expense
	Balance             interfaces.Balance
	Unrecognizeddeposit interfaces.Unrecognizeddeposit
	Applicationpayment  interfaces.Applicationpayment
}

type Controller struct {
	Repositories *Repositories
}

func createPostgresqlDatabase(cfg *config.Configuration) *pg.DB {
	db, err := postgresql.NewPosgresqlDB(cfg.Psql.Host, cfg.Psql.Username, cfg.Psql.Username, cfg.Psql.Database, cfg.Environment)
	if err != nil {
		// log.Bg().Error("Database.Error", zap.Error(err))
		panic("database error")
	}
	return db
}

func NewController(cfg *config.Configuration) *Controller {
	db := createPostgresqlDatabase(cfg)

	return &Controller{
		Repositories: &Repositories{
			Account:             postgresql.NewAccount(db),
			Accountant:          postgresql.NewAccountant(db),
			Discount:            postgresql.NewDiscount(db),
			Deposit:             postgresql.NewDeposit(db),
			Provider:            postgresql.NewProvider(db),
			House:               postgresql.NewHouse(db),
			Condo:               postgresql.NewCondo(db),
			Expense:             postgresql.NewExpense(db),
			Balance:             postgresql.NewBalance(db),
			Unrecognizeddeposit: postgresql.NewUnrecognizeddeposit(db),
			Applicationpayment:  postgresql.NewApplicationpayment(db),
		},
	}
}
