package algebra

import "testing"

func TestShouldCalculateAdjoinMatrix2By2(t *testing.T) {

	var isValidOperation = false

	var matrix = CreateMatrix(3, 3)
	matrix[0][0] = 1
	matrix[0][1] = 2
	matrix[0][2] = 3
	matrix[1][0] = 0
	matrix[1][1] = 4
	matrix[1][2] = 5
	matrix[2][0] = 1
	matrix[2][1] = 0
	matrix[2][2] = 6

	adjointMatrix, err := AdjointMatrix(matrix)

	isValidOperation = FloatEquals(adjointMatrix[0][0], 24) &&
		FloatEquals(adjointMatrix[0][1], -12) &&
		FloatEquals(adjointMatrix[0][2], -2) &&
		FloatEquals(adjointMatrix[1][0], 5) &&
		FloatEquals(adjointMatrix[1][1], 3) &&
		FloatEquals(adjointMatrix[1][2], -5) &&
		FloatEquals(adjointMatrix[2][0], -4) &&
		FloatEquals(adjointMatrix[2][1], 2) &&
		FloatEquals(adjointMatrix[2][2], 4) &&
		err == nil

	if !isValidOperation {
		t.Error("Failed to calculate Adjoint matrix !")
	}
}
