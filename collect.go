package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/pterm/pterm"
	"go.bug.st/serial"
)

func collectData(device string, s serial.Port, sample_size int, interval_value int) {
	file_name, csv_name := makeFileName(device, sample_size, interval_value)
	block_size := sample_size / 8
	buff := make([]byte, block_size)
	num_loop := 1
	total_bytes := 0
	time.Sleep(100 * time.Millisecond)
	for {
		start := time.Now()
		total_bytes += block_size
		// Reads up to 100 bytes
		n, err := s.Read(buff)
		if err != nil {
			log.Fatalf("error while reading buff: %s", err)
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
		fmt.Printf(pterm.LightCyan("Collecting data - Loop: %d - Total bytes collected: %d - "), num_loop, total_bytes)
		fmt.Printf(pterm.LightGreen("Number of \"1\"s: %v\n"), ones)

		// Open a csvfile to write the data to

		t := time.Now()
		time_str := fmt.Sprintf("%02v:%02v:%02v", t.Hour(),
			t.Minute(),
			t.Second())
		ones_str := fmt.Sprintf("%v", ones)
		record := [][]string{{time_str + " " + ones_str}}
		csvfile, err := os.OpenFile(csv_name, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		defer csvfile.Close()

		recordWriter := csv.NewWriter(csvfile)

		for _, value := range record {
			err := recordWriter.Write(value)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
		}
		recordWriter.Flush()

		num_loop += 1
		// sleep for x second
		elapsed := time.Since(start)
		//fmt.Printf("Time elapsed: %s\n", elapsed)
		time_to_sleep := (time.Duration(interval_value) * time.Second) - elapsed
		//fmt.Printf("Sleeping for: %v seconds \n", time_to_sleep)
		time.Sleep(time_to_sleep)
	}
}
