package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"go.bug.st/serial"
	"go.bug.st/serial/enumerator"
)

/*

https://github.com/bugst/go-serial

*/

// serialMonitorCmd represents the serialMonitor command
var serialMonitorCmd = &cobra.Command{
	Use:   "monitor",
	Short: "show the monitor serial USB",
	Run: func(cmd *cobra.Command, args []string) {
		baudrate, _ := cmd.Flags().GetInt("baudrate")
		portName, _ := cmd.Flags().GetString("port")

		if portName == "" {
			portName, _ = getFisrtPort()
		}

		fmt.Println("serialMonitor called", portName, baudrate)

		mode := &serial.Mode{
			BaudRate: baudrate,
			DataBits: 8,
			Parity:   serial.NoParity,
			StopBits: serial.OneStopBit,
		}

		port, err := serial.Open(portName, mode)
		if err != nil {
			log.Fatal(err)
		}

		buf := make([]byte, 256)
		for {
			n, err := port.Read(buf)
			if err != nil {
				log.Fatal(err)
			}
			if n == 0 {
				fmt.Println("\nEOF")
				break
			}

			fmt.Printf("%s", buf[:n])
		}
	},
}

func init() {
	serialCmd.AddCommand(serialMonitorCmd)

	serialMonitorCmd.Flags().StringP("port", "p", "", "set port serial")
	serialMonitorCmd.Flags().IntP("baudrate", "b", 115200, "set baudrate")

	serialMonitorCmd.Flags().BoolP("follow", "f", false, "reconnect enable")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serialMonitorCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
}

func getFisrtPort() (string, bool) {
	ports, err := enumerator.GetDetailedPortsList()

	if err != nil {
		log.Fatal(err)
	}

	if len(ports) == 0 {
		log.Fatal("No serial ports found!")
		return "", false

	}

	return ports[0].Name, true
}
