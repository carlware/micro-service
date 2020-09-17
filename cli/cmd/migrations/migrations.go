package migrations

import (
	"arquil/accounts/cli/config"
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/carlware/go-common/log"
	"go.uber.org/zap"

	"github.com/pressly/goose"
	"github.com/spf13/cobra"
)

var cfgFilePath string

const dbType = "postgres"

var MigrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Perform database migrations",
	Args:  cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		config.InitConfig(cfgFilePath)
		config.InitLogger(ctx, config.Conf)
		conf := config.Conf

		hostAndPort := strings.Split(conf.Psql.Host, ":")
		dbConf := fmt.Sprintf("user=%s dbname=%s sslmode=disable password=%s host=%s port=%s", conf.Psql.Username, conf.Psql.Database, conf.Psql.Password, hostAndPort[0], hostAndPort[1])

		err := goose.SetDialect(dbType)
		if err != nil {
			log.Bg().Fatal("goose: failed to select DB type: %v\n", zap.Error(err))
		}

		db, err := sql.Open(dbType, dbConf)
		if err != nil {
			log.Bg().Fatal("goose: failed to connect DB type: %v\n", zap.Error(err))
		}

		if len(args) < 1 {
			log.Bg().Fatal("Command is not passed")
		}

		command := args[0]

		if err := goose.Run(command, db, "migrations", args[1:]...); err != nil {
			log.Sugar().Fatalf("goose %v: %v", command, err)
		}
	},
}

func init() {
	MigrateCmd.Flags().StringVarP(&cfgFilePath, "config", "c", "", "Configuration File")
}
