package dispatchers

import (
	"carlware/accounts/cli/config"
	"carlware/accounts/internal/interfaces"
	"carlware/accounts/internal/interfaces/postgresql"
	"carlware/accounts/internal/interfaces/memorydb"

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

func NewPostgrestController(cfg *config.Configuration) *Controller {
	db := createPostgresqlDatabase(cfg)

	return &Controller{
		Repositories: &Repositories{
			Account:             postgresql.NewAccount(db),
		},
	}
}

func NewMemorydbController() *Controller {

	return &Controller{
		Repositories: &Repositories{
			Account:             memorydb.NewAccount(),
		},
	}
}

func NewController(cfg *config.Configuration) *Controller {

	switch(cfg.Database){
	case "postgresql":
		return NewPostgrestController(cfg)
	case "memorydb":
		return NewMemorydbController()
	default:
		return nil
	}
}
