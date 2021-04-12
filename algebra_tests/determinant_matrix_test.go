package algebra_tests

import (
	"github.com/ajantonov/machine-learning/algebra"
	"testing"
)

func TestShouldCalculate2by2MatrixDeterminantSuccessfully(t *testing.T) {

	var matrix = algebra.CreateMatrix(2, 2)
	matrix[0][0] = 1
	matrix[0][1] = 2
	matrix[1][0] = 3
	matrix[1][1] = 4

	var determinant = algebra.Determinant(matrix, len(matrix))
	if !algebra.FloatEquals(determinant, -2) {
		t.Error("Failed to calculate Determinant for 2x2 matrix expected : -2 get : ", determinant)
	}
}

func TestShouldCalculate3by3MatrixDeterminantSuccessfully(t *testing.T) {

	var matrix = algebra.CreateMatrix(3, 3)
	matrix[0][0] = 1
	matrix[0][1] = 2
	matrix[0][2] = 3
	matrix[1][0] = 4
	matrix[1][1] = 5
	matrix[1][2] = 6
	matrix[2][0] = 7
	matrix[2][1] = 8
	matrix[2][2] = 9

	var determinant = algebra.Determinant(matrix, len(matrix))
	if !algebra.FloatEquals(determinant, 0) {
		t.Error("Failed to calculate Determinant for 3x3 matrix expected : -2 get : ", determinant)
	}
}
