package cmd

import (
	"github.com/spf13/cobra"
)

func NewSerialCmd() *cobra.Command {
	serialCmd := &cobra.Command{
		Use:   "serial",
		Short: "Simple serial interface to your devices",
	}

	return serialCmd
}
