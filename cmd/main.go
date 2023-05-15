package main

import (
	"fmt"
	"os"

	"PlatformScienceExercise/internal/calculate"
	"PlatformScienceExercise/internal/io"
	"PlatformScienceExercise/internal/matrices"

	hungarianAlgorithm "github.com/oddg/hungarian-algorithm"
)

func main() {
	// checks for proper amount of arguments
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run main.go <shipment_file> <driver_file>")
		os.Exit(1)
	}

	// Parse arguments and make 2 lists of strings
	shipmentFile := os.Args[1]
	driverFile := os.Args[2]
	shipmentLines, driverLines := io.ParseFiles(shipmentFile, driverFile)

	// Create a matrix of Suitability Scores and prep another for hungarian algo
	matrixA, matrixB := matrices.MakeMatrices(shipmentLines, driverLines)

	// Find the optimal pairings of shipments/drivers
	answerKey, err := hungarianAlgorithm.Solve(matrixB)
	if err != nil {
		errMsg := fmt.Sprintf("Error solving matrixB: %s\nMatrixB: %v", err, matrixB)
		fmt.Println(errMsg)
	}

	// Calculate Suitability Score and print the correct pairings
	results := calculate.FinalResults(shipmentLines, driverLines, matrixA, answerKey)

	// write the results to a textfile
	writeErr := os.WriteFile("output.txt", []byte(results), 0644)
	if writeErr != nil {
		fmt.Println("Error writing file", writeErr)
	}
}
