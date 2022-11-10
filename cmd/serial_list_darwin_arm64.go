package cmd

import (
	"fmt"
	"log"

	"go.bug.st/serial"
)

func serialPortsList() {
	ports, err := serial.GetPortsList()
	if err != nil {
		log.Fatal(err)
	}
	if len(ports) == 0 {
		log.Fatal("No serial ports found!")
	}
	for _, port := range ports {
		fmt.Printf("Found port: %s\n", port)
	}
}
