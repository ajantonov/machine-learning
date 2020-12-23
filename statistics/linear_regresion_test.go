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

func TestLinearRegressionShouldReturnNonZeroAlphaCoefficientAndZeroBetaCoefficient(t *testing.T) {

	alpha, beta, err := LinearRegression([]float64{1, 2, 3}, []float64{3, 6, 9})

	if !FloatEqualsWithEpsilon(alpha, float64(3), 0.0001) ||
		!FloatEqualsWithEpsilon(beta, float64(0), 0.0001) ||
		err != nil {
		t.Errorf("Failed to calculate alpha = 3 and beta = 0 for linear regression alpha result = %f beta result = %f ", alpha, beta)
	}
}

func TestLinearRegressionShouldReturnZeroForAlphaAndNonZeroBetaCoefficient(t *testing.T) {

	alpha, beta, err := LinearRegression([]float64{1, 2, 3}, []float64{3, 3, 3})

	if !FloatEqualsWithEpsilon(alpha, float64(0), 0.0001) ||
		!FloatEqualsWithEpsilon(beta, float64(3), 0.0001) ||
		err != nil {
		t.Errorf("Failed to calculate alpha = 0 and beta = 3 for linear regression alpha result = %f beta result = %f ", alpha, beta)
	}
}

/*
Add test for different size of inputs and outputs

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
