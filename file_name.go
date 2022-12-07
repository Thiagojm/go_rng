package main

import (
	"fmt"
	"time"

	"github.com/pterm/pterm"
)

// function to make a file name based on the current time and date
func makeFileName(device string, sample_size int, interval_value int) (string, string) {
	t := time.Now()
	file_name := fmt.Sprintf("%d%02d%02d-%02d%02d%02d_%s_s%d_i%d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), device, sample_size, interval_value)
	fmt.Printf(pterm.LightRed("Using file name: %v\n"), file_name)
	bin_file := file_name + ".bin"
	csv_file := file_name + ".csv"
	return bin_file, csv_file
}
