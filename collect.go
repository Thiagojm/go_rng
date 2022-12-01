package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"go.bug.st/serial"
)

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
		for i := range buff {
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
		//fmt.Printf("Time elapsed: %s\n", elapsed)
		time_to_sleep := (time.Duration(interval_value) * time.Second) - elapsed
		time.Sleep(time_to_sleep)
	}
}
