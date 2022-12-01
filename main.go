package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"go.bug.st/serial"
)

// function to make a file name based on the current time and date
func makeFileName(sample_size, interval_value int) string {
	t := time.Now()
	return fmt.Sprintf("%d%02d%02d-%02d%02d%02d_trng_s%d_i%d.bin", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), sample_size, interval_value)
}

func collectData(s serial.Port, sample_size int, interval_value int) {
	file_name := makeFileName(sample_size, interval_value)
	block_size := sample_size / 8
	buff := make([]byte, block_size)
	num_loop := 1
	total_bytes := 0
	for {
		start := time.Now()
		total_bytes += block_size
		// Reads up to 100 bytes
		n, err := s.Read(buff)
		if err != nil {
			log.Fatal(err)
		}
		if n == 0 {
			fmt.Println("\nEOF")
		}

		// Open a file to write the data to
		file, err := os.OpenFile(file_name, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

		if err != nil {
			fmt.Printf("Could not open %s\n", file_name)
			return
		}

		defer file.Close()

		_, err2 := file.Write(buff)

		if err2 != nil {
			fmt.Printf("Could not write text to %s\n", file_name)
		}

		// buff to string
		var sb strings.Builder
		for i, _ := range buff {
			sb.WriteString(fmt.Sprintf("%08b", buff[i]))
		}
		binString := sb.String()
		// Counting "1"s
		ones := strings.Count(binString, "1")
		fmt.Printf("Collecting data - Loop: %d - Total bytes collected: %d - ", num_loop, total_bytes)
		fmt.Printf("Number of \"1\"s: %v\n", ones)

		num_loop += 1
		// sleep for 1 second
		elapsed := time.Since(start)
		time_to_sleep := (time.Duration(interval_value) * time.Second) - elapsed
		time.Sleep(time_to_sleep)
	}
}

func main() {

	// Retrieve the port list
	ports, err := serial.GetPortsList()
	if err != nil {
		log.Fatal(err)
		return
	}
	if len(ports) == 0 {
		log.Fatal("No serial ports found!")
		return
	}

	// Print the list of detected ports
	for _, port := range ports {
		fmt.Printf("Found port: %v\n", port)
	}

	// Open the first serial port detected at 9600bps N81
	mode := &serial.Mode{
		BaudRate: 300,
		Parity:   serial.NoParity,
		DataBits: 8,
		StopBits: serial.OneStopBit,
	}
	port, err := serial.Open(ports[0], mode)
	if err != nil {
		log.Fatal(err)
	}
	sample_size, interval_value := 2048, 1
	collectData(port, sample_size, interval_value)

}
