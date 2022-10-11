package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:     "trading-bot-crawler",
	Short:   "trading-bot-crawler service",
	Long:    "trading-bot-crawler service",
	Version: "0.0.0",
}

// SetVersion inject version from git
func SetVersion(r string) {
	if len(r) > 0 {
		RootCmd.Version = r
	}
	viper.SetDefault("service_version", RootCmd.Version)
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is config.toml)")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.SetConfigName("config.loc")
	}

	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetConfigType("toml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Cannot read config file: %s", err)
	}

	fmt.Println("Using config file:", viper.ConfigFileUsed())
}
