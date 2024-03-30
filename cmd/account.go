package cmd

import (
	"github.com/spf13/cobra"

	address_util "telegraph/pkg/crypto_utils/address"
)

var accountCmd = &cobra.Command{
	Use:   "account",
	Short: "Manage validator public/private keys",
	Long:  `This command oversees the secure generation/storage of a node's private keys`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			switch args[0] {
			case "new":
				generateAccount()
			case "destroy":
				destroyAccount()
			default:
				getAccount()
			}
		} else {
			getAccount()
		}
	},
}

func init() {
	rootCmd.AddCommand(accountCmd)
}

func getAccount() {

}

func generateAccount() {
	address_util.GetAddress()
}

func destroyAccount() {

}
