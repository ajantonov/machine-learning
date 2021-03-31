package main

import (
	"fmt"
	"github.com/ajantonov/machine-learning/algebra"
)

func main() {

	matrixX := algebra.CreateMatrix(1, 2)
	matrixX[0][0] = 4.0
	matrixX[0][1] = 5

	transposedMatrixX, err := algebra.TransposeMatrix(matrixX)
	if err != nil {
		fmt.Println("Failed to transpose matrix")
	}
	algebra.PrintMatrix(transposedMatrixX)

	resultMatrix, err := algebra.MultiplyMatrices(transposedMatrixX, matrixX)
	algebra.PrintMatrix(resultMatrix)
}
