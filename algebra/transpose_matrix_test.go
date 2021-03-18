package algebra

import "testing"

func TestShouldTransposeMatrix2By2(t *testing.T) {

	isValidOperation := false

	var matrix = CreateMatrix(2, 3)
	matrix[0][0] = 1
	matrix[0][1] = 2
	matrix[0][2] = 3
	matrix[1][0] = 4
	matrix[1][1] = 5
	matrix[1][2] = 6

	transposedMatrix, err := TransposeMatrix(matrix)

	isValidOperation = FloatEquals(transposedMatrix[0][0], 1) &&
		FloatEquals(transposedMatrix[0][1], 4) &&
		FloatEquals(transposedMatrix[1][0], 2) &&
		FloatEquals(transposedMatrix[1][1], 5) &&
		FloatEquals(transposedMatrix[2][0], 3) &&
		FloatEquals(transposedMatrix[2][1], 6) &&
		err == nil

	if !isValidOperation {
		t.Error("Failed to transpose 2x2 matrix")
	}
}

func TestShouldTransposeMatrix3By3(t *testing.T) {
	isValidOperation := false

	var matrix = CreateMatrix(3, 3)
	matrix[0][0] = 1
	matrix[0][1] = 2
	matrix[0][2] = 3
	matrix[1][0] = 4
	matrix[1][1] = 5
	matrix[1][2] = 6
	matrix[2][0] = 7
	matrix[2][1] = 8
	matrix[2][2] = 9

	transposedMatrix, err := TransposeMatrix(matrix)

	isValidOperation = FloatEquals(transposedMatrix[0][0], 1) &&
		FloatEquals(transposedMatrix[0][1], 4) &&
		FloatEquals(transposedMatrix[0][2], 7) &&
		FloatEquals(transposedMatrix[1][0], 2) &&
		FloatEquals(transposedMatrix[1][1], 5) &&
		FloatEquals(transposedMatrix[1][2], 8) &&
		FloatEquals(transposedMatrix[2][0], 3) &&
		FloatEquals(transposedMatrix[2][1], 6) &&
		FloatEquals(transposedMatrix[2][2], 9) &&
		err == nil

	if !isValidOperation {
		t.Error("Failed to transpose 2x2 matrix")
	}
}
