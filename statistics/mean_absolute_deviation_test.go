package statistics

import (
	"testing"
)

/**
MeanAbsoluteDeviation Unit Tests
1. Should return error when array is empty  -  TestMeanAbsoluteDeviationShouldReturnErrorWhenSamplesIsZero
2. Should return result 1.33 when array has elements { 1, 4, 4 } - TestMeanShouldReturnOneWhenArrayContainsOne
**/
func TestMeanAbsoluteDeviationShouldReturnErrorWhenSamplesIsZero(t *testing.T) {

	samples := []float64{1, 3, 5}
	mean := float64(3)

	_, err := MeanAbsoluteDeviation(samples, mean)

	if err != nil {
		t.Error("Failed to return error when array length is zero! ")
	}
}

func TestMeanAbsoluteDeviationShouldReturnValidResult(t *testing.T) {

	samples := []float64{1, 4, 4}
	mean := float64(3)

	result, _ := MeanAbsoluteDeviation(samples, mean)

	if FloatEquals(result, 1.33) {
		t.Errorf("Failed to return valid result : %f !", result)
	}
}
