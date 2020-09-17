package cmd

import (
	"arquil/accounts/cli/cmd/migrations"
	"arquil/accounts/cli/cmd/serve"
	"arquil/accounts/cli/cmd/version"

	"github.com/spf13/cobra"
)

// RootCmd describes root command of the tool
var mainCmd = &cobra.Command{
	Use:   "condo",
	Short: "Microservice to manage condo administrative tasks.",
}

func init() {
	mainCmd.AddCommand(version.VersionCmd)
	mainCmd.AddCommand(serve.ServerCmd)
	mainCmd.AddCommand(migrations.MigrateCmd)
}

// Execute main command
func Execute() error {
	return mainCmd.Execute()
}
