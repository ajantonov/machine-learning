package algebra_tests

import (
	"github.com/ajantonov/machine-learning/algebra"
	"testing"
)

func TestAddMatrixOn2By2ShouldBeSuccessful(t *testing.T) {
	t.Parallel()

	matrixA := [][]float64{{1, 1}, {1, 1}}

	matrixB := [][]float64{{2, 2}, {2, 2}}

	matrixC := algebra.AddMatrices(matrixA, matrixB)

	validAddition := true
	for i := 0; i < len(matrixC); i++ {
		for j := 0; j < len(matrixC[i]); j++ {
			if matrixC[i][j] != 3 {
				validAddition = false
				break
			}
		}
	}

	if !validAddition {
		t.Error("Failed to add matrices 2x2")
	}
}

func TestAddMatrixOn3By3ShouldBeSuccessful(t *testing.T) {
	t.Parallel()

	matrixA := [][]float64{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}}

	matrixB := [][]float64{{2, 2, 2}, {2, 2, 2}, {2, 2, 2}}

	matrixC := algebra.AddMatrices(matrixA, matrixB)

	validAddition := true
	for i := 0; i < len(matrixC); i++ {
		for j := 0; j < len(matrixC[i]); j++ {
			if matrixC[i][j] != 3 {
				validAddition = false
				break
			}
		}
	}

	if !validAddition {
		t.Error("Failed to add matrices 2x2")
	}
}
