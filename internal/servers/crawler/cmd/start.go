package cmd

import (
	"github.com/spf13/cobra"

	"github.com/anvh2/trading-bot/internal/servers/crawler/server"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start tranding-bot-crawler service",
	Long:  "Start tranding-bot-crawler service",
	RunE: func(cmd *cobra.Command, args []string) error {
		server := server.New()
		return server.Start()
	},
}

func init() {
	RootCmd.AddCommand(startCmd)
}
