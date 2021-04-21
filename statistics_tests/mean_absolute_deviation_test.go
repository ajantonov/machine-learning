package statistics_tests

import (
	"github.com/ajantonov/machine-learning/statistics"
	"testing"
)

/**
MeanAbsoluteDeviation Unit Tests
1. Should return error when array is empty  -  TestMeanAbsoluteDeviationShouldReturnErrorWhenSamplesIsZero
2. Should return result 1.33 when array has elements { 1, 4, 4 } - TestMeanAbsoluteDeviationShouldReturnValidResult
**/
func TestMeanAbsoluteDeviationShouldReturnErrorWhenSamplesIsZero(t *testing.T) {

	t.Parallel()

	samples := []float64{1, 3, 5}
	mean := float64(3)

	_, err := statistics.MeanAbsoluteDeviation(samples, mean)

	if err != nil {
		t.Error("Failed to return error when array length is zero! ")
	}
}

func TestMeanAbsoluteDeviationShouldReturnValidResult(t *testing.T) {

	t.Parallel()

	samples := []float64{1, 4, 4}
	mean := float64(3)

	result, _ := statistics.MeanAbsoluteDeviation(samples, mean)

	if !statistics.FloatEqualsWithEpsilon(result, 1.33, 0.004) {
		t.Errorf("Failed to return valid result! Expected value 1.33 result : %f !", result)
	}
}
