package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"go.bug.st/serial"
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
		follow, _ := cmd.Flags().GetBool("follow")

		if portName == "" {
			// TODO() add
			if follow && false {
				stopSpinner := showBusySpinner("Polling for available serial device...")
				for {
					break
				}
				stopSpinner()
			}
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

	// TODO() falta
	serialMonitorCmd.Flags().BoolP("follow", "f", false, "reconnect enable")
}

func getFisrtPort() (string, bool) {
	ports, err := serial.GetPortsList()

	if err != nil {
		log.Fatal(err)
	}

	if len(ports) == 0 {
		log.Fatal("No serial ports found!") // TODO() mover
		return "", false

	}

	return ports[0], true
}
