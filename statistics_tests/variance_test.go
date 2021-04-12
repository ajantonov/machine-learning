package statistics_tests

import (
	"github.com/ajantonov/machine-learning/statistics"
	"testing"
)

/**
Variance Unit Tests
1. Should return error when array is empty  - TestVarianceShouldFailWhenHasSmallerOrEqualThanOneSample
2. Should return result zero when array has { 1, 1 } elements - TestVarianceShouldReturnValidResponseForTwoSamples
3. Should return result 33.2 when array has { 17, 15, 23, 7, 9, 13 } elements - TestVarianceShouldReturnValidResponseForMoreSamples
**/
func TestVarianceShouldFailWhenHasSmallerOrEqualThanOneSample(t *testing.T) {

	var mean float64
	var samples = []float64{1}

	_, err := statistics.Variance(samples, mean)

	if err == nil {
		t.Errorf("Variance failed to return error when has one sample!")
	}
}

func TestVarianceShouldReturnValidResponseForTwoSamples(t *testing.T) {

	var mean float64 = 1
	var samples = []float64{1, 1}

	result, _ := statistics.Variance(samples, mean)

	if !statistics.FloatEquals(float64(0), result) {
		t.Errorf("Variance failed to calculate variance! Expected 0 returned %f for { 1, 1 }!", result)
	}
}

func TestVarianceShouldReturnValidResponseForMoreSamples(t *testing.T) {

	var samples = []float64{17, 15, 23, 7, 9, 13}
	var mean = float64(14)

	result, _ := statistics.Variance(samples, mean)

	if !statistics.FloatEquals(33.2, result) {
		t.Errorf("Function Variance failed to calculate variance! Expected 0 returned %f for { 1, 2, 3, 4, 5, 6 }!", result)
	}
}
