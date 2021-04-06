package main

import (
	"fmt"
	"github.com/ajantonov/machine-learning/algebra"
)

func main() {

	calculatedCoefficientsAreRight := false
	betaTwo := float64(0)
	betaOne := float64(0)
	betaZero := float64(0)

	matrixX := algebra.CreateMatrix(6, 3)
	matrixX[0][0] = 1
	matrixX[0][1] = 7
	matrixX[0][2] = 560
	matrixX[1][0] = 1
	matrixX[1][1] = 3
	matrixX[1][2] = 220
	matrixX[2][0] = 1
	matrixX[2][1] = 3
	matrixX[2][2] = 340
	matrixX[3][0] = 1
	matrixX[3][1] = 4
	matrixX[3][2] = 80
	matrixX[4][0] = 1
	matrixX[4][1] = 6
	matrixX[4][2] = 150
	matrixX[5][0] = 1
	matrixX[5][1] = 7
	matrixX[5][2] = 330

	algebra.PrintMatrix(matrixX)

	// Step 1 : Подготвяне на матрица matrixY
	matrixY := algebra.CreateMatrix(6, 1)
	matrixY[0][0] = 16.68
	matrixY[1][0] = 11.50
	matrixY[2][0] = 12.03
	matrixY[3][0] = 14.88
	matrixY[4][0] = 13.75
	matrixY[5][0] = 18.11
	algebra.PrintMatrix(matrixY)

	transposedMatrixX, err := algebra.TransposeMatrix(matrixX)
	if err != nil {
		fmt.Println("Failed to transpose matrix!")
	}

	transposedMatrixXByMatrixX, err := algebra.MultiplyMatrices(transposedMatrixX, matrixX)
	if err != nil {
		fmt.Println("Failed to multiply transposes matrixX by matrixX !")
	}

	inverseMatrixOfTransposedMatrixXByMatrixX, err := algebra.InverseMatrix(transposedMatrixXByMatrixX)
	if err != nil {
		fmt.Println("Failed to invert transposes matrixX by matrixX !")
	}

	hatMatrix, err := algebra.MultiplyMatrices(inverseMatrixOfTransposedMatrixXByMatrixX, transposedMatrixX)
	if err != nil {
		fmt.Println("Failed to multiply hat matrix by transposes matrixX !")
	}

	finalResultMatrix, err := algebra.MultiplyMatrices(hatMatrix, matrixY)
	if err != nil {
		fmt.Println("Failed to multiply hat matrix by matrixY !")
	}

	betaTwo = finalResultMatrix[0][0]
	betaOne = finalResultMatrix[0][1]
	betaZero = finalResultMatrix[0][2]

	calculatedCoefficientsAreRight = algebra.FloatEqualsWithEpsilon(2.34123115, betaTwo, 0.1) &&
		algebra.FloatEqualsWithEpsilon(1.61590721, betaOne, 0.1) &&
		algebra.FloatEqualsWithEpsilon(0.01438483, betaZero, 0.1)

	if !calculatedCoefficientsAreRight {
		fmt.Printf("Failed to calculate predictors %f :  %f : %f expected 7 : 10 : 50 ",
			betaTwo, betaOne, betaZero)
	}
}
