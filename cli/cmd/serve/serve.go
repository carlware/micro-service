package serve

import (
	"carlware/accounts/cli/config"
	"carlware/accounts/cli/dispatchers"
	"carlware/accounts/cli/dispatchers/graphql"
	"context"

	"github.com/carlware/go-common/log"
	"github.com/oklog/run"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var cfgFilePath string

var ServerCmd = &cobra.Command{
	Use:     "serve",
	Aliases: []string{"s"},
	Short:   "Starts services",
	Run: func(cmd *cobra.Command, args []string) {
		var group run.Group
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		config.InitConfig(cfgFilePath)
		config.InitLogger(ctx, config.Conf)

		// Starting banner
		log.For(ctx).Info("Starting kingdom services ...")

		ctrl := dispatchers.NewController(config.Conf)

		// Add to goroutine group
		group.Add(
			func() error {
				log.For(ctx).Info("GraphQL server listening ...", zap.String("address", config.Conf.GraphQL.Port))
				return graphql.NewGraphQL(ctx, config.Conf, ctrl)
			},
			func(e error) {
				// log.For(ctx).Info("Shutting GRPC server down")
				// graphql.GracefulStop()
			},
		)

		if err := group.Run(); err != nil {
			log.For(ctx).Fatal("Unable to start services", zap.Error(err))
		}
	},
}

func init() {
	ServerCmd.Flags().StringVarP(&cfgFilePath, "config", "c", "", "Configuration File")
}
