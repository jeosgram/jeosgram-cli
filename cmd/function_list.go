package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// functionListCmd represents the functionList command
var functionListCmd = &cobra.Command{
	Use:   "list",
	Short: "Show functions provided by your device(s)",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("functionList called")
	},
}

func init() {
	functionCmd.AddCommand(functionListCmd)
}
