package main

import (
	"log"

	"go.bug.st/serial"
)

func main() {

	// Retrieve the port list
	// ports, err := serial.GetPortsList()
	// if err != nil {
	// 	log.Fatal(err)
	// 	return
	// }
	// fmt.Println(ports)
	// if len(ports) == 0 {
	// 	log.Fatal("No serial ports found!")
	// 	return
	// }

	// // Print the list of detected ports
	// for _, port := range ports {
	// 	fmt.Printf("Found port: %v\n", port)
	// }
	p := get_correct_port()
	// Open the first serial port detected at 9600bps N81
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
