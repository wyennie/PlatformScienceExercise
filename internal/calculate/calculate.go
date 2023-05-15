package calculate

import (
	"fmt"
	"unicode"
)

// Counts the vowels in a string and returns that number x 1.5
func evenBase(driverName string, vowelMap map[rune]bool) float64 {
	vowelCount := 0
	for _, char := range driverName {
		if vowelMap[unicode.ToLower(char)] {
			vowelCount++
		}
	}

	return float64(vowelCount) * 1.5
}

// Counts the consonants in a string and returns that number
func oddBase(driverName string, vowelMap map[rune]bool) float64 {
	consonantCount := 0
	for _, char := range driverName {
		if !vowelMap[unicode.ToLower(char)] && unicode.IsLetter(char) {
			consonantCount++
		}
	}

	return float64(consonantCount)
}

// Checks if the length of two strings have common factors other than 1. If they
// do, returns baseSS * 1.5, otherwise returns baseSS
func commonFactors(driverName string, shipmentDestination string, baseSS float64) float64 {
	driverNameLength := len([]rune(driverName))
	destLength := len([]rune(shipmentDestination))

	for i := 2; i <= driverNameLength && i <= destLength; i++ {
		if driverNameLength%i == 0 && destLength%i == 0 {
			return baseSS * 1.5
		}
	}

	return baseSS
}

func SuitabilityScore(shipmentDestination string, driverName string) float64 {
	vowelMap := map[rune]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true}

	// Calculate the base SS
	var baseSS float64
	if len(shipmentDestination)%2 == 0 {
		baseSS = evenBase(driverName, vowelMap)
	} else {
		baseSS = oddBase(driverName, vowelMap)
	}

	// Check for common factors
	baseSS = commonFactors(driverName, shipmentDestination, baseSS)
	return baseSS
}

func FinalResults(shipmentLines []string, driverLines []string, matrixA [][]float64, answerKey []int) string {
	var suitabilityScore float64
	var results string

	for i := range answerKey {
		column := answerKey[i]
		row := i
		suitabilityScore += matrixA[row][column]
		shipment := shipmentLines[row]
		driver := driverLines[column]

		results += fmt.Sprintf("\nShipment Address: %s\t Driver Name: %s", shipment, driver)
	}

	results += fmt.Sprintf("\n\nTotal Suitability Score: %v", suitabilityScore)
	fmt.Println(results)

	return results
}
