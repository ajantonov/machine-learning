package statistics_tests

import (
	"github.com/ajantonov/machine-learning/statistics"
	"testing"
)

func TestShouldReturnErrorWhenInputsParameterIsSetToZero(t *testing.T) {

	t.Parallel()

	_, _, err := statistics.LinearRegression(nil, []float64{0})

	if err != nil && err.Error() != "invalid parameter x" {
		t.Error("Failed to return error when inputs parameter is with size zero!")
	}
}

func TestShouldReturnErrorWhenOutputsParameterIsSetToZero(t *testing.T) {

	t.Parallel()

	_, _, err := statistics.LinearRegression([]float64{0}, nil)

	if err != nil && err.Error() != "invalid parameter y" {
		t.Error("Failed to return error when y parameter is with size zero!")
	}
}

func TestShouldReturnErrorWhenSizeOfInputsAndOutputsIsNotEqual(t *testing.T) {

	t.Parallel()

	_, _, err := statistics.LinearRegression([]float64{1}, []float64{1, 2})

	if err != nil && err.Error() != "different size of x and y" {
		t.Error("Failed to validate x and y parameters!")
	}
}

func TestLinearRegressionShouldReturnNonZeroAlphaCoefficientAndZeroBetaCoefficient(t *testing.T) {

	t.Parallel()

	slope, intercept, err := statistics.LinearRegression([]float64{1, 2, 3}, []float64{3, 6, 9})

	if !statistics.FloatEqualsWithEpsilon(slope, float64(3), 0.0001) ||
		!statistics.FloatEqualsWithEpsilon(intercept, float64(0), 0.0001) ||
		err != nil {
		t.Errorf("Failed to calculate slope = 3 and intercept = 0 for linear regression slope result = %f intercept result = %f ", slope, intercept)
	}
}

func TestLinearRegressionShouldReturnZeroForAlphaAndNonZeroBetaCoefficient(t *testing.T) {

	t.Parallel()

	slope, intercept, err := statistics.LinearRegression([]float64{1, 2, 3}, []float64{3, 3, 3})

	if !statistics.FloatEqualsWithEpsilon(slope, float64(0), 0.0001) ||
		!statistics.FloatEqualsWithEpsilon(intercept, float64(3), 0.0001) ||
		err != nil {
		t.Errorf("Failed to calculate slope = 0 and intercept = 3 for linear regression slope result = %f intercept result = %f ", slope, intercept)
	}
}

func TestLinearRegressionShouldReturnNonZeroAlphaAndBetaCoefficients(t *testing.T) {

	t.Parallel()

	x := []float64{23, 26, 30, 34, 43, 48, 52, 57, 58}
	y := []float64{651, 762, 856, 1063, 1190, 1298, 1421, 1440, 1518}

	slope, intercept, err := statistics.LinearRegression(x, y)

	if !statistics.FloatEqualsWithEpsilon(slope, 23.4227, 0.0001) ||
		!statistics.FloatEqualsWithEpsilon(intercept, 167.6829, 0.0001) ||
		err != nil {
		t.Errorf("Failed to calculate slope = 23.4227 and intercept = 167.6829 for linear regression slope result = %f intercept result = %f ", slope, intercept)
	}
}
