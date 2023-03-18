package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// functionListCmd represents the functionList command
func NewFunctionListCmd() *cobra.Command {

	functionListCmd := &cobra.Command{
		Use:   "list",
		Short: "Show functions provided by your device(s)",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("functionList called")
		},
	}
	return functionListCmd
}
