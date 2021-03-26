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

	matrixA := [][]float64{{1, 2}, {3, 4}}
	matrixB := [][]float64{{1, 2}, {3, 4}}

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

func TestShouldMultiplyNonSquareMatricesCorrectly(t *testing.T) {

	matrixX := CreateMatrix(6, 3)
	matrixX[0][0] = 1
	matrixX[0][1] = 7
	matrixX[0][2] = 560
	matrixX[1][0] = 1
	matrixX[1][1] = 3
	matrixX[1][2] = 220
	matrixX[2][0] = 1
	matrixX[2][1] = 3
	matrixX[2][2] = 340
	matrixX[3][0] = 1
	matrixX[3][1] = 4
	matrixX[3][2] = 80
	matrixX[4][0] = 1
	matrixX[4][1] = 6
	matrixX[4][2] = 150
	matrixX[5][0] = 1
	matrixX[5][1] = 7
	matrixX[5][2] = 330

	transposedMatrixX := CreateMatrix(3, 6)
	transposedMatrixX[0][0] = 1
	transposedMatrixX[0][1] = 1
	transposedMatrixX[0][2] = 1
	transposedMatrixX[0][3] = 1
	transposedMatrixX[0][4] = 1
	transposedMatrixX[0][5] = 1
	transposedMatrixX[1][0] = 7
	transposedMatrixX[1][1] = 3
	transposedMatrixX[1][2] = 3
	transposedMatrixX[1][3] = 4
	transposedMatrixX[1][4] = 6
	transposedMatrixX[1][5] = 7
	transposedMatrixX[2][0] = 560
	transposedMatrixX[2][1] = 220
	transposedMatrixX[2][2] = 340
	transposedMatrixX[2][3] = 80
	transposedMatrixX[2][4] = 150
	transposedMatrixX[2][5] = 330

	matrixResult, err := MultiplyMatrices(transposedMatrixX, matrixX)

	if err != nil {
		t.Error("Failed to multiply matrices : ", err)
	}

	if !FloatEquals(matrixResult[0][0], 6) ||
		!FloatEquals(matrixResult[0][1], 30) ||
		!FloatEquals(matrixResult[0][2], 1680) ||
		!FloatEquals(matrixResult[1][0], 30) ||
		!FloatEquals(matrixResult[1][1], 168) ||
		!FloatEquals(matrixResult[1][2], 9130) ||
		!FloatEquals(matrixResult[2][0], 1680) ||
		!FloatEquals(matrixResult[2][1], 9130) ||
		!FloatEquals(matrixResult[2][2], 615400) {
		t.Error("Failed to multiply two matrices with size 2x2!")
	}
}
