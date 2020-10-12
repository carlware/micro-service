package version

import (
	"carlware/accounts/internal/version"
	"fmt"

	"github.com/spf13/cobra"
)

var displayFull bool

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display service version",
	Run: func(cmd *cobra.Command, args []string) {
		if displayFull {
			fmt.Printf("%s", version.Full())
		} else {
			fmt.Printf("%s", version.Ver())
		}
	},
}

func init() {
	VersionCmd.Flags().BoolVar(&displayFull, "json", false, "Display build info as json")
}
