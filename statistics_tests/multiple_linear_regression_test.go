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
	Test with multiple regression model with two predictors : y = 50 + 10x1 + 7x2
*/
func TestShouldCalculateMultipleRegressionWithTwoPredictors(t *testing.T) {

	matrixX := algebra.CreateMatrix(5, 3)
	matrixX[0][0] = 0
	matrixX[0][1] = 0
	matrixX[0][2] = 1

	matrixX[1][0] = 1
	matrixX[1][1] = 0
	matrixX[1][2] = 1

	matrixX[2][0] = 0
	matrixX[2][1] = 1
	matrixX[2][2] = 1

	matrixX[3][0] = 1
	matrixX[3][1] = 1
	matrixX[3][2] = 1

	matrixX[4][0] = 2
	matrixX[4][1] = 1
	matrixX[4][2] = 1

	matrixY := algebra.CreateMatrix(5, 1)
	matrixY[0][0] = 7
	matrixY[1][0] = 57
	matrixY[2][0] = 17
	matrixY[3][0] = 67
	matrixY[4][0] = 117

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
/*
	Test with multiple regression model : y = - 2.68 + 9.5x
*/
