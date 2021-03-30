package main

import (
	"fmt"
	"github.com/ajantonov/machine-learning/algebra"
)

func main() {

	matrixX := algebra.CreateMatrix(1, 7)
	matrixX[0][0] = 4.0
	matrixX[0][1] = 4.5
	matrixX[0][2] = 5.0
	matrixX[0][3] = 5.5
	matrixX[0][4] = 6.0
	matrixX[0][5] = 6.5
	matrixX[0][6] = 7.0

	transposedMatrixX, err := algebra.TransposeMatrix(matrixX)
	if err != nil {
		fmt.Println("Failed to transpose matrix")
	}
	algebra.PrintMatrix(transposedMatrixX)

	resultMatrix, _ := algebra.MultiplyMatrices(transposedMatrixX, matrixX)
	algebra.PrintMatrix(resultMatrix)

}
