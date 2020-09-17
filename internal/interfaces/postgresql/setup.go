package postgresql

import (
	"condomilux/condo-admin/internal/models"

	"github.com/carlware/go-common/db/postgresql"

	"github.com/carlware/go-common/errors"
	"github.com/go-pg/pg"
)

// NewPosgresqlDB is...
func NewPosgresqlDB(host, username, password, dbName, env string) (*pg.DB, error) {
	db, err := postgresql.New(host, username, password, dbName)
	if err != nil {
		// fmt.Errorf(err)
		return nil, errors.New(errors.Internal, "Postgresql.Connect", "", err)
	}

	// Create database schema
	dbModels := []interface{}{
		&models.Account{},
	}

	//dropExistingTables := false
	//
	//if env == "local" || env == "test" {
	//	dropExistingTables = true
	//}

	err = postgresql.CreateSchema(db, dbModels, false)
	if err != nil {
		// fmt.Errorf(err)
		return nil, err
	}
	return db, nil
}
