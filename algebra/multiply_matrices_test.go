package algebra

import (
	"github.com/ajantonov/machine-learning/algebra"
	"testing"
)

func TestShouldRequiresValidSizeOfMatrices(t *testing.T) {

	matrixOne := algebra.CreateMatrix(0, 0)
	matrixTwo := algebra.CreateMatrix(0, 0)

	_, err := algebra.MultiplyMatrices(matrixOne, matrixTwo)

	if err == nil {
		t.Error("Invalid validation of input matrices for multiplication!")
	}
}

func TestShouldMultiplyMatricesWithOneRowAndOneColumnCorrectly(t *testing.T) {

	matrixX := algebra.CreateMatrix(1, 1)
	matrixX[0][0] = 4.0

	transposedMatrixX, err := algebra.TransposeMatrix(matrixX)
	if err != nil {
		t.Error("Failed to transpose matrix")
	}

	resultMatrix, _ := algebra.MultiplyMatrices(transposedMatrixX, matrixX)

	isNotResultValid := resultMatrix[0][0] != 16
	matrixHasNotOneRow := len(resultMatrix) != 1
	matrixHasNotOneColumn := len(resultMatrix[0]) != 1

	if isNotResultValid || matrixHasNotOneRow || matrixHasNotOneColumn {
		t.Errorf("Failed to calculate %f ", resultMatrix[0][0])
	}

}

func TestShouldMultiplyMatricesWithOneRowAndTwoColumnsCorrectly(t *testing.T) {

	matrixX := algebra.CreateMatrix(1, 2)
	matrixX[0][0] = 4.0
	matrixX[0][1] = 5.0

	transposedMatrixX, err := algebra.TransposeMatrix(matrixX)
	if err != nil {
		t.Error("Failed to transpose matrix")
	}

	resultMatrix, _ := algebra.MultiplyMatrices(transposedMatrixX, matrixX)

	isNotResultValid := resultMatrix[0][0] != 16 ||
		resultMatrix[0][1] != 20 ||
		resultMatrix[1][0] != 20 ||
		resultMatrix[1][1] != 25
	matrixHasNotOneRow := len(resultMatrix) != 2
	matrixHasNotOneColumn := len(resultMatrix[0]) != 2

	if isNotResultValid || matrixHasNotOneRow || matrixHasNotOneColumn {
		t.Errorf("Failed to calculate multiplication of matrix with one row and two columns!")
	}

}

func TestShouldMultiplyMatricesWithOneRowAndThreeColumnsCorrectly(t *testing.T) {

	matrixX := algebra.CreateMatrix(1, 3)
	matrixX[0][0] = 4.0
	matrixX[0][1] = 5.0
	matrixX[0][2] = 6.0

	transposedMatrixX, err := algebra.TransposeMatrix(matrixX)
	if err != nil {
		t.Error("Failed to transpose matrix")
	}

	resultMatrix, _ := algebra.MultiplyMatrices(transposedMatrixX, matrixX)

	isNotResultValid := resultMatrix[0][0] != 16 ||
		resultMatrix[0][1] != 20 ||
		resultMatrix[0][2] != 24 ||
		resultMatrix[1][0] != 20 ||
		resultMatrix[1][1] != 25 ||
		resultMatrix[1][2] != 30 ||
		resultMatrix[2][0] != 24 ||
		resultMatrix[2][1] != 30 ||
		resultMatrix[2][2] != 36

	matrixHasNotOneRow := len(resultMatrix) != 3
	matrixHasNotOneColumn := len(resultMatrix[0]) != 3

	if isNotResultValid || matrixHasNotOneRow || matrixHasNotOneColumn {
		t.Errorf("Failed to calculate matrices with one row and three columns correctly! Valid Result %v Has Rows %v Has Columns %v",
			isNotResultValid, matrixHasNotOneRow, matrixHasNotOneColumn)
	}

}

func TestShouldMultiply3By3With6By3MatricesCorrectly(t *testing.T) {

	matrixOne := algebra.CreateMatrix(6, 3)
	matrixOne[0][0] = 1
	matrixOne[0][1] = 7
	matrixOne[0][2] = 560
	matrixOne[1][0] = 1
	matrixOne[1][1] = 3
	matrixOne[1][2] = 220
	matrixOne[2][0] = 1
	matrixOne[2][1] = 3
	matrixOne[2][2] = 340
	matrixOne[3][0] = 1
	matrixOne[3][1] = 4
	matrixOne[3][2] = 80
	matrixOne[4][0] = 1
	matrixOne[4][1] = 6
	matrixOne[4][2] = 150
	matrixOne[5][0] = 1
	matrixOne[5][1] = 7
	matrixOne[5][2] = 330

	matrixTwo := algebra.CreateMatrix(3, 3)
	matrixTwo[0][0] = 1.607233
	matrixTwo[0][1] = -0.250638
	matrixTwo[0][2] = -0.000669
	matrixTwo[1][0] = -0.250638
	matrixTwo[1][1] = 0.069809
	matrixTwo[1][2] = -0.000351
	matrixTwo[2][0] = -0.000669
	matrixTwo[2][1] = -0.000351
	matrixTwo[2][2] = 0.000009

	resultMatrix, err := algebra.MultiplyMatrices(matrixOne, matrixTwo)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	algebra.PrintMatrix(resultMatrix)

	isNotResultValid := !algebra.FloatEquals(resultMatrix[0][0], -0.521873) ||
		!algebra.FloatEquals(resultMatrix[0][1], 0.041465) ||
		!algebra.FloatEquals(resultMatrix[0][2], 0.001914) ||
		!algebra.FloatEquals(resultMatrix[1][0], 0.708139) ||
		!algebra.FloatEquals(resultMatrix[1][1], -0.118431) ||
		!algebra.FloatEquals(resultMatrix[1][2], 0.000258) ||
		!algebra.FloatEquals(resultMatrix[2][0], 0.627859) ||
		!algebra.FloatEquals(resultMatrix[2][1], -0.160551) ||
		!algebra.FloatEquals(resultMatrix[2][2], 0.001338) ||
		!algebra.FloatEquals(resultMatrix[3][0], 0.551161) ||
		!algebra.FloatEquals(resultMatrix[3][1], 0.000518) ||
		!algebra.FloatEquals(resultMatrix[3][2], -0.001353) ||
		!algebra.FloatEquals(resultMatrix[4][0], 0.003055) ||
		!algebra.FloatEquals(resultMatrix[4][1], 0.115566) ||
		!algebra.FloatEquals(resultMatrix[4][2], -0.001425) ||
		!algebra.FloatEquals(resultMatrix[5][0], -0.368003) ||
		!algebra.FloatEquals(resultMatrix[5][1], 0.122195) ||
		!algebra.FloatEquals(resultMatrix[5][2], -0.000156)

	matrixHasNotThreeRows := len(resultMatrix) != 6
	matrixHasNotThreeColumns := len(resultMatrix[0]) != 3

	if isNotResultValid || matrixHasNotThreeRows || matrixHasNotThreeColumns {
		t.Errorf("Failed to calculate matrices with one row and three columns correctly! Valid Result %v Has Rows %v Has Columns %v",
			isNotResultValid, matrixHasNotThreeRows, matrixHasNotThreeColumns)
	}

}
func TestShouldNotMultiplyMatrixWithDifferentLengthOfRowAndColumn(t *testing.T) {

	matrixOne := algebra.CreateMatrix(1, 7)
	matrixTwo := algebra.CreateMatrix(7, 2)

	_, err := algebra.MultiplyMatrices(matrixOne, matrixTwo)

	if err != nil {
		t.Error(err.Error())
	}
}

func TestShouldMultiply2by2MatricesCorrectlyWithTheSameNumbers(t *testing.T) {

	multiplicationIsRight := true

	matrixA := [][]float64{{2, 2}, {2, 2}}
	matrixB := [][]float64{{2, 2}, {2, 2}}

	matrixResult, err := algebra.MultiplyMatrices(matrixA, matrixB)

	if err != nil {
		t.Error("Failed to multiply matrices : ", err)
	}

	for i, row := range matrixResult {
		for j := range row {
			if !algebra.FloatEquals(matrixResult[i][j], 8) {
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

	matrixResult, err := algebra.MultiplyMatrices(matrixA, matrixB)

	if err != nil {
		t.Error("Failed to multiply matrices : ", err)
	}

	if !algebra.FloatEquals(matrixResult[0][0], 7) ||
		!algebra.FloatEquals(matrixResult[0][1], 10) ||
		!algebra.FloatEquals(matrixResult[1][0], 15) ||
		!algebra.FloatEquals(matrixResult[1][1], 22) {
		t.Error("Failed to multiply two matrices with size 2x2!")
	}
}

func TestShouldMultiplyNonSquareMatricesCorrectly(t *testing.T) {

	matrixX := algebra.CreateMatrix(6, 3)
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

	transposedMatrixX := algebra.CreateMatrix(3, 6)
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

	matrixResult, err := algebra.MultiplyMatrices(transposedMatrixX, matrixX)

	if err != nil {
		t.Error("Failed to multiply matrices : ", err)
	}

	if !algebra.FloatEquals(matrixResult[0][0], 6) ||
		!algebra.FloatEquals(matrixResult[0][1], 30) ||
		!algebra.FloatEquals(matrixResult[0][2], 1680) ||
		!algebra.FloatEquals(matrixResult[1][0], 30) ||
		!algebra.FloatEquals(matrixResult[1][1], 168) ||
		!algebra.FloatEquals(matrixResult[1][2], 9130) ||
		!algebra.FloatEquals(matrixResult[2][0], 1680) ||
		!algebra.FloatEquals(matrixResult[2][1], 9130) ||
		!algebra.FloatEquals(matrixResult[2][2], 615400) {
		t.Error("Failed to multiply two matrices with size 2x2!")
	}
}
