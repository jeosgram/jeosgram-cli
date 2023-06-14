package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// serialListCmd represents the serialList command
func NewSerialListCmd() *cobra.Command {
	serialListCmd := &cobra.Command{
		Use:   "list",
		Short: "list port USB",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("serialList called")

			serialPortsList()
		},
	}

	return serialListCmd
}
