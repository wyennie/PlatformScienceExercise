package calculate

import "testing"

func TestSuitabilityScore(t *testing.T) {
	// Test case 1
	driverName1 := "Ava Blevins"
	shipmentDestination1 := "272 W. San Juan Dr."
	var expectedScore1 float64 = 6

	score1 := SuitabilityScore(shipmentDestination1, driverName1)

	if score1 != expectedScore1 {
		t.Errorf("SuitabilityScore(%s, %s) = %f; want %f", shipmentDestination1, driverName1, score1, expectedScore1)
	}

	// Test case 2
	driverName2 := "Jaxson Griffith"
	shipmentDestination2 := "19 E. Lakeview Drive"
	var expectedScore2 float64 = 9

	score2 := SuitabilityScore(shipmentDestination2, driverName2)

	if score2 != expectedScore2 {
		t.Errorf("SuitabilityScore(%s, %s) = %f; want %f", shipmentDestination2, driverName2, score2, expectedScore2)
	}

	// Test case 3
	driverName3 := "Marlon Navarro"
	shipmentDestination3 := "7847 Acacia Ave."
	var expectedScore3 float64 = 11.25

	score3 := SuitabilityScore(shipmentDestination3, driverName3)

	if score3 != expectedScore3 {
		t.Errorf("SuitabilityScore(%s, %s) = %f; want %f", shipmentDestination3, driverName3, score3, expectedScore3)
	}
}

func TestFinalResults(t *testing.T) {

	shipmentLines := []string{"272 W. San Juan Dr.", "19 E. Lakeview Drive", "7847 Acacia Ave."}
	driverLines := []string{"Ava Blevins", "Jaxson Griffith", "Marlon Navarro"}

	matrixA := [][]float64{
		{6, 10, 8},
		{6, 9, 11.25},
		{6, 6, 11.25},
	}
	answerKey := []int{1, 2, 0}

	expectedOutput := "\nShipment Address: 272 W. San Juan Dr.\t Driver Name: Jaxson Griffith" +
		"\nShipment Address: 19 E. Lakeview Drive\t Driver Name: Marlon Navarro" +
		"\nShipment Address: 7847 Acacia Ave.\t Driver Name: Ava Blevins" +
		"\n\nTotal Suitability Score: 27.25"

	if output := FinalResults(shipmentLines, driverLines, matrixA, answerKey); output != expectedOutput {
		t.Errorf("FinalResults returned unexpected output: %v", output)
	}
}
