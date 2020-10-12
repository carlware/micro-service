package cmd

import (
	"carlware/accounts/cli/cmd/migrations"
	"carlware/accounts/cli/cmd/serve"
	"carlware/accounts/cli/cmd/version"

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
