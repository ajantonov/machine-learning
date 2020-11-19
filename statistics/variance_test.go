package statistics

import (
	"testing"
)

/**
Variance Unit Tests
1. Should return error when array is empty  - TestVarianceShouldFailWhenHasSmallerOrEqualThanOneSample
2. Should return result zero when array has { 1, 1 } elements - TestVarianceShouldReturnValidResponseForTwoSamples
3. Should return result zero when array has { 1, 2, 3, 4, 5 } elements - TestVarianceShouldReturnValidResponseForMoreSamples
**/
func TestVarianceShouldFailWhenHasSmallerOrEqualThanOneSample(t *testing.T) {

	var mean float64
	var samples = []float64{1}

	_, err := Variance(samples, mean)

	if err == nil {
		t.Errorf("Variance failed to return error when has one sample!")
	}
}

func TestVarianceShouldReturnValidResponseForTwoSamples(t *testing.T) {

	var mean float64 = 1
	var samples = []float64{1, 1}

	result, _ := Variance(samples, mean)

	if !FloatEquals(float64(0), result) {
		t.Errorf("Variance failed to calculate variance! Expected 0 returned %f for { 1, 1 }!", result)
	}
}

func TestVarianceShouldReturnValidResponseForMoreSamples(t *testing.T) {

	var samples = []float64{1, 2, 3, 4, 5, 6}
	var mean = 3.5

	result, _ := Variance(samples, mean)

	if !FloatEquals(float64(0), result) {
		t.Errorf("Function Variance failed to calculate variance! Expected 0 returned %f for { 1, 2, 3, 4, 5 }!", result)
	}
}
