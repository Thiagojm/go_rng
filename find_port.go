package main

import (
	"fmt"
	"log"
	"strings"

	"go.bug.st/serial"
	"go.bug.st/serial/enumerator"
)

func get_correct_port() string {
	ports, err := enumerator.GetDetailedPortsList()
	if err != nil {
		log.Fatal(err)
	}
	if len(ports) == 0 {
		log.Fatal(err)
	}
	for _, port := range ports {
		//fmt.Printf("Port: %s\n", port.Name)
		if port.Product != "" {
			//fmt.Println(port)
			//fmt.Println(port.Product)
			if strings.Contains(port.Product, "TrueRNG") {
				fmt.Printf("Found TrueRNG on %v\n", port.Name)
				p := string(port.Name)
				return p
			}
		}
	}
	ports2, err2 := serial.GetPortsList()
	if err2 != nil {
		log.Fatal(err)
	}
	if len(ports2) == 0 {
		log.Fatal(err)
	}
	return ports2[0]

}
