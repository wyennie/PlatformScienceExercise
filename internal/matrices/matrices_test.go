package matrices

import (
	"testing"
)

func TestMakeMatrices(t *testing.T) {
	shipments := []string{"272 W. San Juan Dr.", "19 E. Lakeview Drive", "7847 Acacia Ave."}
	drivers := []string{"Ava Blevins", "Jaxson Griffith", "Marlon Navarro"}

	expectedMatrixA := [][]float64{
		{6, 10, 8},
		{6, 9, 11.25},
		{6, 6, 11.25},
	}

	expectedMatrixB := [][]int{
		{99399, 98999, 99199},
		{99399, 99099, 98874},
		{99399, 99399, 98874},
	}

	matrixA, matrixB := MakeMatrices(shipments, drivers)

	// Test if the dimensions of the matrices are correct
	if len(matrixA) != len(shipments) || len(matrixA[0]) != len(drivers) {
		t.Errorf("Expected matrixA dimensions (%d, %d), but got (%d, %d)", len(shipments), len(drivers), len(matrixA), len(matrixA[0]))
	}
	if len(matrixB) != len(shipments) || len(matrixB[0]) != len(drivers) {
		t.Errorf("Expected matrixB dimensions (%d, %d), but got (%d, %d)", len(shipments), len(drivers), len(matrixB), len(matrixB[0]))
	}

	// Test if the values in the matrices are correct
	for i := range matrixA {
		for j := range matrixA[i] {
			if matrixA[i][j] != expectedMatrixA[i][j] {
				t.Errorf("Expected matrixA value %f at (%d, %d), but got %f", expectedMatrixA[i][j], i, j, matrixA[i][j])
			}
			if matrixB[i][j] != expectedMatrixB[i][j] {
				t.Errorf("Expected matrixB value %d at (%d, %d), but got %d", expectedMatrixB[i][j], i, j, matrixB[i][j])
			}
		}
	}
}
