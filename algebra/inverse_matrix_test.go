package algebra

import (
	"testing"
)

func TestShouldInvert2By2MatrixCorrectly(t *testing.T) {

	isValidOperation := false

	var matrix = CreateMatrix(2, 2)
	matrix[0][0] = 4
	matrix[0][1] = 3
	matrix[1][0] = 3
	matrix[1][1] = 2

	invertedMatrix, err := InverseMatrix(matrix)

	if err != nil {
		t.Error("Failed to invert matrix : ", err)
	}

	isValidOperation = FloatEquals(invertedMatrix[0][0], -2) &&
		FloatEquals(invertedMatrix[0][1], 3) &&
		FloatEquals(invertedMatrix[1][0], 3) &&
		FloatEquals(invertedMatrix[1][1], -4) &&
		err == nil

	if !isValidOperation {
		t.Error("Failed to invert matrix!")
	}
}

func TestShouldInvert3By3MatrixCorrectly(t *testing.T) {

	isValidOperation := false

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

	invertedMatrix, err := InverseMatrix(matrix)

	if err != nil {
		t.Error("Failed to invert matrix : ", err)
	}

	isValidOperation = FloatEquals(invertedMatrix[0][0], 1.090909) &&
		FloatEquals(invertedMatrix[0][1], -0.545455) &&
		FloatEquals(invertedMatrix[0][2], -0.090909) &&
		FloatEquals(invertedMatrix[1][0], 0.227273) &&
		FloatEquals(invertedMatrix[1][1], 0.136364) &&
		FloatEquals(invertedMatrix[1][2], -0.227273) &&
		FloatEquals(invertedMatrix[2][0], -0.181818) &&
		FloatEquals(invertedMatrix[2][1], 0.090909) &&
		FloatEquals(invertedMatrix[2][2], 0.181818) &&
		err == nil

	PrintMatrix(invertedMatrix)

	if !isValidOperation {
		t.Error("Failed to invert matrix!")
	}
}
