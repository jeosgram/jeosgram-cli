package cmd

import (
	"github.com/spf13/cobra"
)

// serialCmd represents the serial command
var serialCmd = &cobra.Command{
	Use:   "serial",
	Short: "Simple serial interface to your devices",
}

func init() {
	rootCmd.AddCommand(serialCmd)
}
