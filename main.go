package main

import (
	"fmt"

	"go.bug.st/serial"
)

func main() {
	sample_size, interval_value := 2048, 1
	// Retrieve the port list
	p := get_correct_port()

	// Open the first serial port detected at 300bps N81
	mode := &serial.Mode{
		BaudRate: 300,
		Parity:   serial.NoParity,
		DataBits: 8,
		StopBits: serial.OneStopBit,
	}
	port, err := serial.Open(p, mode)
	if err != nil {
		fmt.Println("Error opening port: no Serial found")
		fmt.Println("Starting in Pseudo-Random mode - use for testing only.")
		device := "pseudo"
		pseudo_collect(device, sample_size, interval_value)
	} else {
		device := "trng"
		collectData(device, port, sample_size, interval_value)
	}

}
