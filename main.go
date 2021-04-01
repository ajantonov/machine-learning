package main

import (
	"fmt"
	"github.com/ajantonov/machine-learning/algebra"
)

func main() {

	matrixA := [][]float64{{2, 2}, {2, 2}}
	matrixB := [][]float64{{2, 2}, {2, 2}}

	matrixX, err := algebra.MultiplyMatrices(matrixA, matrixB)

	transposedMatrixX, err := algebra.TransposeMatrix(matrixX)
	if err != nil {
		fmt.Println("Failed to transpose matrix")
	}
	algebra.PrintMatrix(transposedMatrixX)

	resultMatrix, err := algebra.MultiplyMatrices(transposedMatrixX, matrixX)
	algebra.PrintMatrix(resultMatrix)
}
