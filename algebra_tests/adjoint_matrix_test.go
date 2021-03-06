package algebra_tests

import (
	"github.com/ajantonov/machine-learning/algebra"
	"testing"
)

func TestShouldCalculateAdjoinMatrix2By2(t *testing.T) {
	t.Parallel()

	var isValidOperation = false

	var matrix = algebra.CreateMatrix(3, 3)
	matrix[0][0] = 1
	matrix[0][1] = 2
	matrix[0][2] = 3
	matrix[1][0] = 0
	matrix[1][1] = 4
	matrix[1][2] = 5
	matrix[2][0] = 1
	matrix[2][1] = 0
	matrix[2][2] = 6

	adjointMatrix, err := algebra.AdjointMatrix(matrix)

	isValidOperation = algebra.FloatEquals(adjointMatrix[0][0], 24) &&
		algebra.FloatEquals(adjointMatrix[0][1], -12) &&
		algebra.FloatEquals(adjointMatrix[0][2], -2) &&
		algebra.FloatEquals(adjointMatrix[1][0], 5) &&
		algebra.FloatEquals(adjointMatrix[1][1], 3) &&
		algebra.FloatEquals(adjointMatrix[1][2], -5) &&
		algebra.FloatEquals(adjointMatrix[2][0], -4) &&
		algebra.FloatEquals(adjointMatrix[2][1], 2) &&
		algebra.FloatEquals(adjointMatrix[2][2], 4) &&
		err == nil

	if !isValidOperation {
		t.Error("Failed to calculate Adjoint matrix !")
	}
}
