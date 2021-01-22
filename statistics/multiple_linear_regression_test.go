package statistics

import (
	"gonum.org/v1/gonum/mat"
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

	// Step 0 : Подготвяне на матрица X
	// y = 50 + 10x1 + 7x2
	// x0 = { 50, 50, 50, 50, 50 } - bias
	// x1 = {  0,  1,  2,  3,  4 } x 10
	// x2 = {  4,  3,  2,  1,  0 } x 7
	x := []float64{4, 0, 50, 3, 1, 50, 2, 2, 50, 1, 3, 50, 0, 4, 50}
	matX := mat.NewDense(5, 3, x)
	matPrint(matX)

	// Step 1 : Подготвяне на матрица Y
	y := []float64{78, 81, 84, 87, 90}
	matY := mat.NewDense(5, 1, y)
	matPrint(matY)

	// Step 2 : Транспониране на матрица
	// -> X.T()
	matPrint(matX.T())

	// Step 3 : Смятане на обратна матрица
	// func (m *Dense) Inverse(a Matrix) error

	//var matInverted mat.Dense
	//err := matInverted.Inverse(matX)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("%v\n", mat.Formatted(&matInverted, mat.Prefix("")))

	// Step 4 : Умножение на матрици
	// Multiplication is pretty straightforward
	// D := mat.NewDense(3, 3, nil)
	// D.Product(A, B.T())
	// println("A * B'")
	// matPrint(D)
	// Step 5 : Смятане на израза за матрици

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
