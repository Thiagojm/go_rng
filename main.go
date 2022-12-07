package main

import (
	"fmt"
	"time"

	"github.com/pterm/pterm"
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

	fmt.Print(pterm.LightGreen("What's the sample size in bits (please insert a number divisible by 8)? - default = 2048 bits: "))
	input_chan := make(chan int, 1)
	go getInput(input_chan)

	select {
	case i := <-input_chan:
		sample_size = i
		fmt.Printf(pterm.LightYellow("The sample size is %v bits\n"), sample_size)
	case <-time.After(5000 * time.Millisecond):
		fmt.Printf(pterm.LightYellow("\nThe sample size is the default value of %v bits\n"), sample_size)
	}

	fmt.Print(pterm.LightGreen("What the interval in seconds? - default = 1s: "))
	go getInput(input_chan)

	select {
	case i := <-input_chan:
		interval_value = i
		fmt.Printf(pterm.LightYellow("The interval is %v second(s)\n"), interval_value)
	case <-time.After(5000 * time.Millisecond):
		fmt.Printf(pterm.LightYellow("\nThe interval is the default value of %v second(s)\n"), interval_value)
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
