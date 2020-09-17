// +build integration

package postgresql

import (
	"testing"

	"github.com/go-pg/pg"
	"github.com/stretchr/testify/assert"
)

func TestNewPosgresqlDB(t *testing.T) {
	type args struct {
		host     string
		username string
		password string
		dbName   string
		env      string
	}
	tests := []struct {
		name string
		args func(t *testing.T) args

		want1      *pg.DB
		wantErr    bool
		inspectErr func(err error, t *testing.T) // use for more precise error evaluation after test
	}{
		{
			args: func(t *testing.T) args {
				return args{
					host:     dbConfig.host,
					username: dbConfig.username,
					password: dbConfig.password,
					dbName:   dbConfig.dbName,
					env:      dbConfig.env,
				}
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tArgs := tt.args(t)

			got1, err := NewPosgresqlDB(tArgs.host, tArgs.username, tArgs.password, tArgs.dbName, tArgs.env)

			assert.NotNil(t, got1)

			if (err != nil) != tt.wantErr {
				t.Fatalf("NewPosgresqlDB error = %v, wantErr: %t", err, tt.wantErr)
			}

			if tt.inspectErr != nil {
				tt.inspectErr(err, t)
			}
		})
	}
}
