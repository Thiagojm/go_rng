package main

import (
	"log"

	"go.bug.st/serial"
)

func main() {

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
		log.Fatal(err)
	}
	sample_size, interval_value := 2048, 1
	collectData(port, sample_size, interval_value)

}
