package statistics

import (
	"testing"
)

func TestShouldReturnErrorWhenInputsParameterIsSetToZero(t *testing.T) {

	_, _, err := LinearRegression(nil, []float64{0})

	if err.Error() != "invalid parameter inputs" {
		t.Error("Failed to return error when inputs parameter is with size zero!")
	}

}

func TestShouldReturnErrorWhenOutputsParameterIsSetToZero(t *testing.T) {

	_, _, err := LinearRegression([]float64{0}, nil)

	if err.Error() != "invalid parameter outputs" {
		t.Error("Failed to return error when outputs parameter is with size zero!")
	}
}

/*

func TestLinearRegressionShouldReturnNonZeroAlphaCoefficientAndZeroBetaCoefficient(t *testing.T) {

	if !FloatEqualsWithEpsilon(alphaResult, float64(0), 0.0001) &&
		!FloatEqualsWithEpsilon(betaResult, float64(0), 0.0001) {
		t.Errorf("Failed to calculate alpha = 23.4227 and beta = 167.6829 for linear regression alpha result = %f beta result = %f ", alphaResult, betaResult)
	}

}

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
