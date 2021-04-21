package algebra_tests

import (
	"github.com/ajantonov/machine-learning/algebra"
	"testing"
)

func TestShouldTransposeMatrix2By2(t *testing.T) {

	t.Parallel()

	isValidOperation := false

	var matrix = algebra.CreateMatrix(2, 3)
	matrix[0][0] = 1
	matrix[0][1] = 2
	matrix[0][2] = 3
	matrix[1][0] = 4
	matrix[1][1] = 5
	matrix[1][2] = 6

	transposedMatrix, err := algebra.TransposeMatrix(matrix)

	isValidOperation = algebra.FloatEquals(transposedMatrix[0][0], 1) &&
		algebra.FloatEquals(transposedMatrix[0][1], 4) &&
		algebra.FloatEquals(transposedMatrix[1][0], 2) &&
		algebra.FloatEquals(transposedMatrix[1][1], 5) &&
		algebra.FloatEquals(transposedMatrix[2][0], 3) &&
		algebra.FloatEquals(transposedMatrix[2][1], 6) &&
		err == nil

	if !isValidOperation {
		t.Error("Failed to transpose 2x2 matrix")
	}
}

func TestShouldTransposeMatrix3By3(t *testing.T) {

	t.Parallel()

	isValidOperation := false

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

	transposedMatrix, err := algebra.TransposeMatrix(matrix)

	isValidOperation = algebra.FloatEquals(transposedMatrix[0][0], 1) &&
		algebra.FloatEquals(transposedMatrix[0][1], 4) &&
		algebra.FloatEquals(transposedMatrix[0][2], 7) &&
		algebra.FloatEquals(transposedMatrix[1][0], 2) &&
		algebra.FloatEquals(transposedMatrix[1][1], 5) &&
		algebra.FloatEquals(transposedMatrix[1][2], 8) &&
		algebra.FloatEquals(transposedMatrix[2][0], 3) &&
		algebra.FloatEquals(transposedMatrix[2][1], 6) &&
		algebra.FloatEquals(transposedMatrix[2][2], 9) &&
		err == nil

	if !isValidOperation {
		t.Error("Failed to transpose 2x2 matrix")
	}
}
