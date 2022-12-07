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
