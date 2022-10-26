package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// webhookCmd represents the webhook command
var webhookCmd = &cobra.Command{
	Use:   "webhook",
	Short: "Manage webhooks that react to device event streams",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("webhook called")
	},
}

func init() {
	rootCmd.AddCommand(webhookCmd)
}
