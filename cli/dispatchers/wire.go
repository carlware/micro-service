package dispatchers

import (
	"arquil/accounts/cli/config"
	"arquil/accounts/internal/interfaces"
	"arquil/accounts/internal/interfaces/postgresql"

	"github.com/go-pg/pg"
)

type Repositories struct {
	Account             interfaces.Account
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
		},
	}
}
