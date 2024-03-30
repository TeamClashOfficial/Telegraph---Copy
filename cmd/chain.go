package cmd

import (
	"telegraph/log"
	networkmodule "telegraph/modules/network"

	listener "telegraph/pkg/listener"

	diContainer "telegraph/container"

	"github.com/spf13/cobra"
)

var networkServiceInstance networkmodule.NetworkService
var listenerInstance listener.Listener

var container = diContainer.GetContainer().Invoke(func(networkService networkmodule.NetworkService, listener listener.Listener) {
	networkServiceInstance = networkService
	listenerInstance = listener
})

var cmdArgs = networkmodule.Network{
	ChainID: 0,
}

func init() {
	rootCmd.AddCommand(chainCmd)
	chainCmd.Flags().IntVarP(&cmdArgs.ChainID, "id", "", 0, "Chain ID of the referenced blockchain")
	chainCmd.Flags().StringVarP(&cmdArgs.ChainType, "type", "", "", "Chain Type of of the referenced blockchain(commonly 'EVM')")
	chainCmd.Flags().StringVarP(&cmdArgs.Name, "name", "", "", "Chain symbol of the referenced blockchain")
	chainCmd.Flags().StringVarP(&cmdArgs.ContractAddress, "port", "", "", "Port contract address on the referenced blockchain")
	chainCmd.Flags().StringVarP(&cmdArgs.EVMHTTPURL, "evmhttp", "", "", "Provider HTTPS url for referenced blockchain")
	chainCmd.Flags().StringVarP(&cmdArgs.EVMWSSURL, "evmwss", "", "", "Provider WSS url for referenced blockchain")
}

var chainCmd = &cobra.Command{
	Use:     "chain",
	Aliases: []string{"c"},
	Short:   "Manage supported blockchains",
	Long:    `The chain command is used to add or remove a blockchain from the list of observed chains`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			switch args[0] {
			case "add":
				addChain()
			case "remove":
				removeChain()
			case "edit":
				editChain()
			case "clear":
				removeAllChains()
			default:
				getAllChains()
			}
		} else {
			getAllChains()
		}
	},
}

func getAllChains() []networkmodule.Network {
	chains, err := networkServiceInstance.GetNetworks()
	if err != nil {
		log.Fatal("networks.GetNetworks error: ", err)
	}

	return chains
}

func addChain() {
	log.Debug("ChainId: ", cmdArgs.ChainID)
	log.Debug("ChainType: ", cmdArgs.ChainType)
	log.Debug("Name: ", cmdArgs.Name)

	if err := networkServiceInstance.AddNetwork(cmdArgs); err != nil {
		log.Error("Error adding network: ", err)
	}

	listenerInstance.CreateConnection(cmdArgs)
}

func removeChain() {
	if err := networkServiceInstance.RemoveNetwork(cmdArgs.Name); err != nil {
		log.Error("Error removing network: ", err)
	}
}

func editChain() {
	if err := networkServiceInstance.UpdateNetwork(cmdArgs); err != nil {
		log.Error("Error updating network: ", err)
	}
}

func removeAllChains() {
	if err := networkServiceInstance.RemoveAllNetworks(); err != nil {
		log.Error("Error removing all networks: ", err)
	}
}
