package cmd

import (
	"github.com/spf13/cobra"
)

// functionCmd represents the function command
func NewFunctionCmd() *cobra.Command {
	functionCmd := &cobra.Command{
		Use:   "function",
		Short: "Call functions on your device",
	}
	return functionCmd
}
