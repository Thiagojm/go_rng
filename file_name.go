package main

import (
	"fmt"
	"time"
)

// function to make a file name based on the current time and date
func makeFileName(sample_size, interval_value int) (string, string) {
	t := time.Now()
	file_name := fmt.Sprintf("%d%02d%02d-%02d%02d%02d_trng_s%d_i%d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), sample_size, interval_value)
	fmt.Printf("Using file name: %v\n", file_name)
	bin_file := file_name + ".bin"
	csv_file := file_name + ".csv"
	return bin_file, csv_file
}
