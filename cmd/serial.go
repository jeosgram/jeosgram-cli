package cmd

import (
	"github.com/spf13/cobra"
)

// serialCmd represents the serial command
var serialCmd = &cobra.Command{
	Use:   "serial",
	Short: "A brief description of your command",
}

func init() {
	rootCmd.AddCommand(serialCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serialCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serialCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
