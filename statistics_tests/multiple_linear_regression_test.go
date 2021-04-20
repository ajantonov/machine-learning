package statistics_tests

import (
	"github.com/ajantonov/machine-learning/algebra"
	"github.com/ajantonov/machine-learning/statistics"
	"testing"
)

/*
Useful links :
1. http://mezeylab.cb.bscb.cornell.edu/labmembers/documents/supplement%205%20-%20multiple%20regression.pdf
*/

/*
	Test with multiple regression model : y = - 2.68 + 9.5x
*/
func TestShouldCalculateMultipleRegressionWithOnePredictor(t *testing.T) {

	matrixX := algebra.CreateMatrix(5, 2)
	matrixX[0][0] = 1
	matrixX[0][1] = 0

	matrixX[1][0] = 1
	matrixX[1][1] = 1

	matrixX[2][0] = 1
	matrixX[2][1] = 2

	matrixX[3][0] = 1
	matrixX[3][1] = 3

	matrixX[4][0] = 1
	matrixX[4][1] = 4

	matrixY := algebra.CreateMatrix(5, 1)
	matrixY[0][0] = -2.68
	matrixY[1][0] = 6.82
	matrixY[2][0] = 16.32
	matrixY[3][0] = 25.82
	matrixY[4][0] = 35.32

	matrixResult, err := statistics.MultipleRegression(matrixX, matrixY)
	if err != nil {
		t.Error("Failed to calculate beta coefficients ", err.Error())
	}
	algebra.PrintMatrix(matrixResult)

	if !algebra.FloatEquals(matrixResult[0][0], -2.68) ||
		!algebra.FloatEquals(matrixResult[1][0], 9.5) {
		t.Errorf("Failed to calculate multiple regression expected [ 50, 1, 1, 1] result [ %f , %f, %f, %f] ",
			matrixResult[0][0], matrixResult[1][0], matrixResult[2][0], matrixResult[3][0])
	}

}

/*
	Test with multiple regression model with two predictors : y = 50 + 10x1 + 7x2
*/
func TestShouldCalculateMultipleRegressionWithTwoPredictors(t *testing.T) {

	matrixX := algebra.CreateMatrix(5, 3)
	matrixX[0][0] = 1
	matrixX[0][1] = 0
	matrixX[0][2] = 1

	matrixX[1][0] = 1
	matrixX[1][1] = 0
	matrixX[1][2] = 1

	matrixX[2][0] = 1
	matrixX[2][1] = 1
	matrixX[2][2] = 1

	matrixX[3][0] = 1
	matrixX[3][1] = 2
	matrixX[3][2] = 1

	matrixX[4][0] = 1
	matrixX[4][1] = 1
	matrixX[4][2] = 2

	matrixY := algebra.CreateMatrix(5, 1)
	matrixY[0][0] = 57
	matrixY[1][0] = 57
	matrixY[2][0] = 67
	matrixY[3][0] = 77
	matrixY[4][0] = 74

	matrixResult, err := statistics.MultipleRegression(matrixX, matrixY)
	if err != nil {
		t.Error("Failed to calulate beta coefficients ", err.Error())
	}
	algebra.PrintMatrix(matrixResult)

	if !algebra.FloatEquals(matrixResult[0][0], 50) ||
		!algebra.FloatEquals(matrixResult[1][0], 10) ||
		!algebra.FloatEquals(matrixResult[2][0], 7) {
		t.Errorf("Failed to calculate multiple regression expected [ 50, 10, 7] result [ %f , %f, %f] ",
			matrixResult[0][0], matrixResult[1][0], matrixResult[2][0])
	}

}

/*
	Test with multiple regression model : y = 50 + x1 + x2 + x1x2
*/
func TestShouldCalculateMultipleRegressionWithThreePredictors(t *testing.T) {

	matrixX := algebra.CreateMatrix(5, 4)
	matrixX[0][0] = 1
	matrixX[0][1] = 0
	matrixX[0][2] = 0
	matrixX[0][3] = 0

	matrixX[1][0] = 1
	matrixX[1][1] = 0
	matrixX[1][2] = 1
	matrixX[1][3] = 0

	matrixX[2][0] = 1
	matrixX[2][1] = 1
	matrixX[2][2] = 1
	matrixX[2][3] = 1

	matrixX[3][0] = 1
	matrixX[3][1] = 1
	matrixX[3][2] = 2
	matrixX[3][3] = 2

	matrixX[4][0] = 1
	matrixX[4][1] = 2
	matrixX[4][2] = 2
	matrixX[4][3] = 4

	matrixY := algebra.CreateMatrix(5, 1)
	matrixY[0][0] = 50
	matrixY[1][0] = 51
	matrixY[2][0] = 53
	matrixY[3][0] = 55
	matrixY[4][0] = 58

	matrixResult, err := statistics.MultipleRegression(matrixX, matrixY)
	if err != nil {
		t.Error("Failed to calculate beta coefficients ", err.Error())
	}
	algebra.PrintMatrix(matrixResult)

	if !algebra.FloatEquals(matrixResult[0][0], 50) ||
		!algebra.FloatEquals(matrixResult[1][0], 1) ||
		!algebra.FloatEquals(matrixResult[2][0], 1) ||
		!algebra.FloatEquals(matrixResult[2][0], 1) {
		t.Errorf("Failed to calculate multiple regression expected [ 50, 1, 1, 1] result [ %f , %f, %f, %f] ",
			matrixResult[0][0], matrixResult[1][0], matrixResult[2][0], matrixResult[3][0])
	}

}
