package statistics

import (
	"fmt"
	"github.com/ajantonov/machine-learning/algebra"
	"testing"
)

/*
Useful links :
1. https://www.scribbr.com/statistics/multiple-linear-regression/
2. http://mezeylab.cb.bscb.cornell.edu/labmembers/documents/supplement%205%20-%20multiple%20regression.pdf
*/

/*
	Test with multiple regression model with two predictors : y = 50 + 10x1 + 7x2

*/
func TestShouldCalculateMultipleRegressionWithTwoPredictors(t *testing.T) {
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
		fmt.Println("Error ", err)
		panic(err)
	}
	algebra.PrintMatrix(transposedMatrixX)

	multipliedMatricesX, err := algebra.MultiplyMatrices(transposedMatrixX, matrixX)
	if err != nil {
		fmt.Println("Error ", err)
		panic(err)
	}
	algebra.PrintMatrix(multipliedMatricesX)

	calculatedCoefficientsAreRight = FloatEqualsWithEpsilon(7, betaTwo, 0.1) &&
		FloatEqualsWithEpsilon(10, betaOne, 0.1) &&
		FloatEqualsWithEpsilon(50, betaZero, 0.1)

	if !calculatedCoefficientsAreRight {
		t.Errorf("Failed to calculate predictors %f :  %f : %f expected 7 : 10 : 50 ",
			betaTwo, betaOne, betaZero)
	}
}

/*
	Test with multiple regression model : y = 50 + x1 + x2 + x1x2

*/
