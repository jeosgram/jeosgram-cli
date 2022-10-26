package cmd

import (
	"github.com/spf13/cobra"
)

// variableCmd represents the variable command
var variableCmd = &cobra.Command{
	Use:   "variable",
	Short: "A brief description of your command",
}

func init() {
	rootCmd.AddCommand(variableCmd)
}
