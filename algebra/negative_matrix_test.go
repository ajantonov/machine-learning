package algebra

import "testing"

func TestShouldReturnNegativeMatrix2By2(t *testing.T) {

	isNegative := true

	matrix := [][]float64{{1, 1}, {1, 1}}

	resultMatrix := NegativeMatrix(matrix)
	for i, row := range resultMatrix {
		for j := range row {
			if resultMatrix[i][j] >= 0 {
				isNegative = false
				break
			}
		}
	}

	if !isNegative {
		t.Error("Failed to do negative matrix")
	}
}

func TestShouldReturnNegativeMatrix3By3(t *testing.T) {

	isNegative := true

	matrix := [][]float64{{3, 3, 3}, {3, 3, 3}, {3, 3, 3}}

	resultMatrix := NegativeMatrix(matrix)
	for i, row := range resultMatrix {
		for j := range row {
			if resultMatrix[i][j] >= 0 {
				isNegative = false
				break
			}
		}
	}

	if !isNegative {
		t.Error("Failed to do negative matrix")
	}
}
