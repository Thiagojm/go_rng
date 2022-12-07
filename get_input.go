package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func getInput(input chan int) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()
	result, err := strconv.Atoi(text)
	if err != nil {
		fmt.Println("Error during conversion")
		return
	}

	input <- result
}

// as
// fmt.Print(pterm.LightGreen("What's the sample size in bits (please insert a number divisible by 8)? - default = 2048 bits: "))
// input_chan := make(chan int, 1)
// go getInput(input_chan)

// select {
// case i := <-input_chan:
// 	sample_size := i
// 	fmt.Printf(pterm.LightYellow("The sample size is %v bits\n"), sample_size)
// case <-time.After(4000 * time.Millisecond):
// 	sample_size := 2048
// 	fmt.Printf(pterm.LightYellow("\nThe sample size is the default value of %v bits\n"), sample_size)
// }

// fmt.Print(pterm.LightGreen("What the interval in seconds? - default = 1s: "))
// go getInput(input_chan)

// select {
// case i := <-input_chan:
// 	interval_value := i
// 	fmt.Printf(pterm.LightYellow("The interval is %v second(s)\n"), interval_value)
// case <-time.After(4000 * time.Millisecond):
// 	interval_value := 1
// 	fmt.Printf(pterm.LightYellow("\nThe interval is the default value of %v second(s)\n"), interval_value)
// }
