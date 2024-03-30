package cmd

import (
	"os"
	"telegraph/log"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// This is used for config file.
	cfgFile string
	rootCmd = &cobra.Command{
		Use:   "telegraph",
		Short: "Telegraph: A lightweight, yet powerful crosschain system",
		Long:  `Telegraph: A lightweight, yet powerful crosschain system that detects and transports messages across various chains`,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.deployer.yaml)")
}

func er(msg interface{}) {
	log.Error("Error:", msg)
	os.Exit(1)
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory
		home, err := homedir.Dir()
		if err != nil {
			er(err)
		}
		viper.AddConfigPath(home)
		viper.SetConfigName(".telegraph")
	}
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err == nil {
		log.Error("Using config file:", viper.ConfigFileUsed())
	}
}
