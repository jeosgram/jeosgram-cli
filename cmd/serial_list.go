package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// serialListCmd represents the serialList command
var serialListCmd = &cobra.Command{
	Use:   "list",
	Short: "list port USB",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("serialList called")

		serialPortsList()
	},
}

func init() {
	serialCmd.AddCommand(serialListCmd)
}
