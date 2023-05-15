package io

import (
	"bufio"
	"fmt"
	"os"
)

// Takes a filename as input and returns a slice of strings & checks for error.
func readLines(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func ParseFiles(shipmentFile string, driverFile string) ([]string, []string) {
	shipmentLines, err := readLines(shipmentFile)
	if err != nil {
		fmt.Printf("Error reading %s: %v\n", shipmentFile, err)
		os.Exit(1)
	}

	driverLines, err := readLines(driverFile)
	if err != nil {
		fmt.Printf("Error reading %s: %v\n", driverFile, err)
		os.Exit(1)
	}

	return shipmentLines, driverLines
}
