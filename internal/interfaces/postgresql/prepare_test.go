// +build integration

package postgresql

import (
	"os"
	"testing"

	"github.com/go-pg/pg"
)

type DBConfig struct {
	host     string
	username string
	password string
	dbName   string
	env      string
}

func getEnv(key, def string) string {
	env := os.Getenv(key)
	if len(env) == 0 {
		return def
	}
	return env
}

var dbConfig = DBConfig{
	host:     getEnv("SQL_DB_HOST", "localhost:5432"),
	username: getEnv("SQL_DB_USERNAME", "golang"),
	password: getEnv("SQL_DB_PASSWORD", ""),
	dbName:   getEnv("SQL_DB_NAME", "condo_test"),
	env:      "test",
}

var db *pg.DB

func TestMain(m *testing.M) {
	var err error
	db, err = NewPosgresqlDB(dbConfig.host, dbConfig.username, dbConfig.password, dbConfig.dbName, dbConfig.env)
	if err != nil {
		panic(err)
	}
	code := m.Run()
	os.Exit(code)
}
