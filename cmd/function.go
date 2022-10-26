package cmd

import (
	"github.com/spf13/cobra"
)

// functionCmd represents the function command
var functionCmd = &cobra.Command{
	Use:   "function",
	Short: "Call functions on your device",
}

func init() {
	rootCmd.AddCommand(functionCmd)
}
