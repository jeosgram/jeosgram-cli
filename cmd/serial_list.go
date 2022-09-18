package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"go.bug.st/serial/enumerator"
)

// serialListCmd represents the serialList command
var serialListCmd = &cobra.Command{
	Use:   "list",
	Short: "list port USB",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("serialList called")

		ports, err := enumerator.GetDetailedPortsList()
		if err != nil {
			log.Fatal(err)
		}
		if len(ports) == 0 {
			log.Fatal("No serial ports found!")
		}
		for _, port := range ports {
			fmt.Printf("Found port: %s\n", port.Name)
			if port.IsUSB {
				fmt.Printf("\tUSB ID     %s:%s\n", port.VID, port.PID)
				fmt.Printf("\tUSB serial %s\n", port.SerialNumber)
			}
		}
	},
}

func init() {
	serialCmd.AddCommand(serialListCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serialListCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serialListCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
