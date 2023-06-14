package cmd

import (
	"github.com/spf13/cobra"
)

func NewVariableCmd() *cobra.Command {
	variableCmd := &cobra.Command{
		Use:   "variable",
		Short: "Retrieve and monitor variables on your device(s)",
	}
	return variableCmd
}
