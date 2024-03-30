package cmd

import (
	"telegraph/log"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version number of Telegraph node",
	Long:  `This command can be used get the version number of this Telegraph node`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Info("Telegraph v0.0.1-alpha")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
