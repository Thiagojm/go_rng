package main

import (
	"fmt"
	"log"
	"runtime"
	"strings"

	"go.bug.st/serial"
	"go.bug.st/serial/enumerator"
)

func get_correct_port() string {
	os := runtime.GOOS

	if os == "windows" {
		ports, err := enumerator.GetDetailedPortsList()
		if err != nil {
			log.Fatal(err)
		} else if len(ports) == 0 {
			log.Fatal(err)
		}
		// Try to find the port with the correct name
		for _, port := range ports {
			//fmt.Printf("Port: %s\n", port.Name)
			if port.Product != "" {
				//fmt.Println(port)
				//fmt.Println(port.Product)
				if strings.Contains(port.Product, "RNG") || strings.Contains(port.Product, "rng") {
					fmt.Printf("Found TrueRNG on %v\n", port.Name)
					p := string(port.Name)
					return p
				}
			}
		}
		return "Failed"
	} else if os == "linux" {

		// If not found, return the first port
		ports2, err2 := serial.GetPortsList()
		if err2 != nil {
			log.Fatal(err2)
		} else if len(ports2) == 0 {
			log.Fatal("No serial port found")
		}
		return ports2[0]

	} else {
		return "Failed"
	}
}
