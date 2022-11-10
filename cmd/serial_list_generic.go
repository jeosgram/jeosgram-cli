package cmd

import (
	"fmt"
	"log"

	"go.bug.st/serial/enumerator"
)

func serialPortsList() {
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
}
