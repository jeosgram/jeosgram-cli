package cmd

import (
	"github.com/spf13/cobra"
)

// variableCmd represents the variable command
var variableCmd = &cobra.Command{
	Use:   "variable",
	Short: "Retrieve and monitor variables on your device(s)",
}

func init() {
	rootCmd.AddCommand(variableCmd)
}
