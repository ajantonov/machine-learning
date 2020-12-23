package statistics

import (
	"testing"
)

func TestShouldReturnErrorWhenInputsParameterIsSetToZero(t *testing.T) {

	_, _, err := LinearRegression(nil, []float64{0})

	if err.Error() != "invalid parameter x" {
		t.Error("Failed to return error when inputs parameter is with size zero!")
	}
}

func TestShouldReturnErrorWhenOutputsParameterIsSetToZero(t *testing.T) {

	_, _, err := LinearRegression([]float64{0}, nil)

	if err.Error() != "invalid parameter y" {
		t.Error("Failed to return error when y parameter is with size zero!")
	}
}

func TestShouldReturnErrorWhenSizeOfInputsAndOutputsIsNotEqual(t *testing.T) {

	_, _, err := LinearRegression([]float64{1}, []float64{1, 2})

	if err.Error() != "different size of x and y" {
		t.Error("Failed to validate x and y parameters!")
	}
}

func TestLinearRegressionShouldReturnNonZeroAlphaCoefficientAndZeroBetaCoefficient(t *testing.T) {

	slope, intercept, err := LinearRegression([]float64{1, 2, 3}, []float64{3, 6, 9})

	if !FloatEqualsWithEpsilon(slope, float64(3), 0.0001) ||
		!FloatEqualsWithEpsilon(intercept, float64(0), 0.0001) ||
		err != nil {
		t.Errorf("Failed to calculate slope = 3 and intercept = 0 for linear regression slope result = %f intercept result = %f ", slope, intercept)
	}
}

func TestLinearRegressionShouldReturnZeroForAlphaAndNonZeroBetaCoefficient(t *testing.T) {

	slope, intercept, err := LinearRegression([]float64{1, 2, 3}, []float64{3, 3, 3})

	if !FloatEqualsWithEpsilon(slope, float64(0), 0.0001) ||
		!FloatEqualsWithEpsilon(intercept, float64(3), 0.0001) ||
		err != nil {
		t.Errorf("Failed to calculate slope = 0 and intercept = 3 for linear regression slope result = %f intercept result = %f ", slope, intercept)
	}
}

/*


func TestLinearRegressionShouldReturnZeroForAlphaAndNonZeroBetaCoefficient(t *testing.T) {

	if !FloatEqualsWithEpsilon(alphaResult, float64(0), 0.0001) &&
		!FloatEqualsWithEpsilon(betaResult, float64(0), 0.0001) {
		t.Errorf("Failed to calculate alpha = 23.4227 and beta = 167.6829 for linear regression alpha result = %f beta result = %f ", alphaResult, betaResult)
	}
}


func TestLinearRegressionShouldReturnNonZeroAlphaAndBetaCoefficients(t *testing.T) {

	inputs := []float64 { 0, 1, 2, 3 }
	outputs := []float64 { 0, 2, 4, 6 }
	alpha, beta, err := LinearRegression(inputs, outputs)

		if !FloatEqualsWithEpsilon(alpha, 23.4227, 0.0001) &&
			!FloatEqualsWithEpsilon(beta, 167.6829, 0.0001) {
			t.Errorf("Failed to calculate alpha = 23.4227 and beta = 167.6829 for linear regression alpha result = %f beta result = %f ", alphaResult, betaResult)
		}

	t.Fail()
}
*/
