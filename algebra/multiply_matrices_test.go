package algebra

import (
	"testing"
)

func TestShouldMultiply2by2MatricesCorrectlyWithTheSameNumbers(t *testing.T) {

	multiplicationIsRight := true

	matrixA := [][]float64{{2, 2}, {2, 2}}
	matrixB := [][]float64{{2, 2}, {2, 2}}

	matrixResult, err := MultiplyMatrices(matrixA, matrixB)

	if err != nil {
		t.Error("Failed to multiply matrices : ", err)
	}

	for i, row := range matrixResult {
		for j := range row {
			if !FloatEquals(matrixResult[i][j], 8) {
				multiplicationIsRight = false
			}
		}
	}

	if !multiplicationIsRight {
		t.Error("Failed to multiply two matrices with size 2x2!")
	}
}

func TestShouldMultiply2by2MatricesCorrectly(t *testing.T) {

	matrixA := [][]float64{{1, 2},
		{3, 4}}
	matrixB := [][]float64{{1, 2},
		{3, 4}}

	matrixResult, err := MultiplyMatrices(matrixA, matrixB)

	if err != nil {
		t.Error("Failed to multiply matrices : ", err)
	}

	if !FloatEquals(matrixResult[0][0], 7) ||
		!FloatEquals(matrixResult[0][1], 10) ||
		!FloatEquals(matrixResult[1][0], 15) ||
		!FloatEquals(matrixResult[1][1], 22) {
		t.Error("Failed to multiply two matrices with size 2x2!")
	}
}
