package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/pterm/pterm"
	"go.bug.st/serial"
)

// Create a function to check errors
func checkErr(e error) {
	if e != nil {
		fmt.Println(e)
	}
}

func main() {
	var sample_size, interval_value int
	// Check if default.txt exists
	if _, err := os.Stat("default.txt"); err == nil {
		// If it does exist, print a message
		pterm.Success.Println("Reading default.txt and setting variables")
		// Open default.txt
		file, err := os.Open("default.txt")
		checkErr(err)
		// Read the first two lines and convert them to integers
		scanner := bufio.NewScanner(file)
		scanner.Scan()
		firstLine := scanner.Text()
		sample_size, err = strconv.Atoi(firstLine)
		checkErr(err)
		if sample_size%8 != 0 || sample_size == 0 {
			pterm.Error.Println("The sample size is not divisible by 8 or invalid. Please change it in default.txt - setting it to 2048 bits")
			sample_size = 2048
		}
		scanner.Scan()
		secondLine := scanner.Text()
		interval_value, err = strconv.Atoi(secondLine)
		if interval_value == 0 {
			pterm.Error.Println("The interval is 0 or invalid. Please change it in default.txt - setting it to 1 second")
			interval_value = 1
		}
		checkErr(err)
		// Close the file
		file.Close()
		pterm.Info.Printf("The sample size is the default value of %v bits\n", sample_size)
		pterm.Info.Printf("The interval is the default value of %v second(s)\n", interval_value)
	} else {
		sample_size, interval_value = 1, 1
		input_chan := make(chan int, 1)
		for sample_size%8 != 0 {
			fmt.Print(pterm.LightGreen("What's the sample size in bits (please insert a number divisible by 8)? - default = 2048 bits: "))
			go getInput(input_chan)

			select {
			case i := <-input_chan:
				sample_size = i
				pterm.Info.Printf("The sample size is %v bits\n", sample_size)
			case <-time.After(7000 * time.Millisecond):
				sample_size = 2048
				fmt.Println("")
				pterm.Info.Printf("The sample size is the default value of %v bits\n", sample_size)
			}
			if sample_size%8 != 0 {
				pterm.Error.Printf("Please insert a number divisible by 8\n")
			}
		}

		fmt.Printf(pterm.LightGreen("What the interval in seconds? - default = %vs: "), interval_value)
		go getInput(input_chan)

		select {
		case i := <-input_chan:
			interval_value = i
			pterm.Info.Printf("The interval is %v second(s)\n", interval_value)
		case <-time.After(7000 * time.Millisecond):
			fmt.Println("")
			pterm.Info.Printf("The interval is the default value of %v second(s)\n", interval_value)
		}
	}

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
		pterm.Error.Printf("Error opening port: no Serial found\n")
		pterm.Warning.Printf("Starting in Pseudo-Random mode - use for testing only.\n")
		device := "pseudo"
		pseudo_collect(device, sample_size, interval_value)
	} else {
		device := "trng"
		collectData(device, port, sample_size, interval_value)
	}

}
