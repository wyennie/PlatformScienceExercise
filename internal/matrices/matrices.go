package matrices

import (
	"PlatformScienceExercise/internal/calculate"
)

const LARGENUM int = 99999

func hungarianMatrix(scores [][]float64) [][]int {
	preppedMatrix := make([][]int, len(scores))

	for i := range scores {
		currentSlice := make([]int, len(scores))

		for j := range scores[i] {
			score := scores[i][j] * 100
			newScore := int(score)
			finalScore := (LARGENUM - newScore)

			currentSlice[j] = finalScore
		}
		preppedMatrix[i] = currentSlice
	}

	return preppedMatrix
}

func suitabilityScoreMatrix(shipments []string, drivers []string) [][]float64 {
	scoreMatrix := make([][]float64, len(shipments))

	for i := range shipments {
		currentSlice := make([]float64, len(drivers))

		for j := range drivers {
			score := calculate.SuitabilityScore(shipments[i], drivers[j])
			currentSlice[j] = score
		}

		scoreMatrix[i] = currentSlice
	}

	return scoreMatrix
}

func MakeMatrices(shipments []string, drivers []string) ([][]float64, [][]int) {
	matrixA := suitabilityScoreMatrix(shipments, drivers)
	matrixB := hungarianMatrix(matrixA)

	return matrixA, matrixB
}
