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
	fmt.Println("Transposed matrix : ")
	algebra.PrintMatrix(transposedMatrixX)

	multipliedMatricesX, err := algebra.MultiplyMatrices(transposedMatrixX, matrixX)
	if err != nil {
		fmt.Println("Error ", err)
		panic(err)
	}
	//algebra.PrintMatrix(multipliedMatricesX)

	inversedMatrixOfX, err := algebra.InverseMatrix(multipliedMatricesX)
	if err != nil {
		fmt.Println("Error of algebra.InverseMatrix ", err)
		panic(err)
	}
	fmt.Println("Inversed matrix : ")
	algebra.PrintMatrix(inversedMatrixOfX)

	finalResultOfMatrixX, err := algebra.MultiplyMatrices(inversedMatrixOfX, transposedMatrixX)
	if err != nil {
		fmt.Println("Error of algebra.InverseMatrix ", err)
		panic(err)
	}
	algebra.PrintMatrix(finalResultOfMatrixX)

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

/*
	Test with multiple regression model : y = - 2.68 + 9.5x
*/
func TestShouldCalculateMultipleRegressionWithOnePredictor(t *testing.T) {

	t.Skip()

	calculatedCoefficientsAreRight := false
	betaOne := float64(0)
	betaZero := float64(0)

	matrixX := algebra.CreateMatrix(1, 7)
	matrixX[0][0] = 4.0
	matrixX[0][1] = 4.5
	matrixX[0][2] = 5.0
	matrixX[0][3] = 5.5
	matrixX[0][4] = 6.0
	matrixX[0][5] = 6.5
	matrixX[0][6] = 7.0

	algebra.PrintMatrix(matrixX)

	// Step 1 : Подготвяне на матрица matrixY
	matrixY := algebra.CreateMatrix(7, 1)
	matrixY[0][0] = 33
	matrixY[1][0] = 42
	matrixY[2][0] = 45
	matrixY[3][0] = 51
	matrixY[4][0] = 53
	matrixY[5][0] = 61
	matrixY[6][0] = 62
	algebra.PrintMatrix(matrixY)

	transposedMatrixX, err := algebra.TransposeMatrix(matrixX)
	if err != nil {
		fmt.Println("Error ", err)
		panic(err)
	}
	fmt.Println("Transposed matrix : ")
	algebra.PrintMatrix(transposedMatrixX)

	multipliedMatricesX, err := algebra.MultiplyMatrices(transposedMatrixX, matrixX)
	if err != nil {
		fmt.Println("Error ", err)
		panic(err)
	}
	fmt.Println("X'X : ")
	algebra.PrintMatrix(multipliedMatricesX)

	inverseMatrix, err := algebra.InverseMatrix(multipliedMatricesX)
	if err != nil {
		fmt.Println("Error ", err)
		panic(err)
	}
	fmt.Println("X^-1 : ")
	algebra.PrintMatrix(inverseMatrix)

	transposedMatrixXBymatrixY, err := algebra.MultiplyMatrices(transposedMatrixX, matrixY)
	if err != nil {
		fmt.Println("Error ", err)
		panic(err)
	}
	fmt.Println("X'Y : ")
	algebra.PrintMatrix(transposedMatrixXBymatrixY)

	resultMatrix, err := algebra.MultiplyMatrices(inverseMatrix, transposedMatrixXBymatrixY)

	betaOne = resultMatrix[0][0]
	betaZero = resultMatrix[0][1]

	calculatedCoefficientsAreRight = FloatEqualsWithEpsilon(9.5, betaOne, 0.1) &&
		FloatEqualsWithEpsilon(-2.68, betaZero, 0.1)

	if !calculatedCoefficientsAreRight {
		t.Errorf("Failed to calculate predictors %f :  %f expected 9.5 : 2.68 ", betaOne, betaZero)
	}
}
