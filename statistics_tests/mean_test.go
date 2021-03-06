package statistics_tests

import (
	"github.com/ajantonov/machine-learning/statistics"
	"testing"
)

/**
Mean Unit Tests
1. Should return error when array is empty  - TestMeanShouldReturnErrorWhenArrayIsEmpty
2. Should return result one when array has one element with value one - TestMeanShouldReturnOneWhenArrayContainsOne
3. Should return 2 when array has elements from 1 to 3 - func TestMeanShouldReturnTwoForRangeBetweenOneToThree
**/

func TestMeanShouldReturnErrorWhenInputArrayIsEmpty(t *testing.T) {

	t.Parallel()

	var values []float64

	_, err := statistics.Mean(values)

	if err == nil {
		t.Error("Mean function should return error!")
	}
}

func TestMeanShouldReturnOneWhenInputArrayContainsValueOne(t *testing.T) {

	t.Parallel()

	values := []float64{1}

	result, _ := statistics.Mean(values)

	if float64(1) != result {
		t.Errorf("Mean function should return 1 when array contains { 1 }. The result is %f ", result)
	}
}

func TestMeanShouldReturnTwoWhenInputArrayContainsRangeFromOneToThree(t *testing.T) {

	t.Parallel()

	values := []float64{1, 2, 3}

	result, _ := statistics.Mean(values)

	if !statistics.FloatEquals(2, result) {
		t.Errorf("Mean function should return result 2, the result is %f", result)
	}
}

func TestScalingDataScalesTheMean(t *testing.T) {

	t.Parallel()

	const scale = float64(2)
	values := []float64{1 * scale, 2 * scale, 3 * scale}

	result, _ := statistics.Mean(values)

	if !statistics.FloatEquals(2*scale, result) {
		t.Errorf("Mean function should return result 4, the result is %f", result)
	}

}

func TestTranslatingDataTranslatesTheMean(t *testing.T) {

	t.Parallel()

	const intercept = float64(2)
	values := []float64{1 + intercept, 2 + intercept, 3 + intercept}

	result, _ := statistics.Mean(values)

	if !statistics.FloatEquals(2+intercept, result) {
		t.Errorf("Mean function should return result 4, the result is %f", result)
	}
}

func TestTheSumOfSignedDifferencesFromTheMeanIsZero(t *testing.T) {

	t.Parallel()

	values := []float64{1, 2, 3}

	mean, _ := statistics.Mean(values)

	var sumOfSignedDifferences = float64(0)

	for _, value := range values {
		sumOfSignedDifferences += value - mean
	}

	if !statistics.FloatEquals(0, sumOfSignedDifferences) {
		t.Errorf("The sum of signed differences from the mean is 0, the result is %f", sumOfSignedDifferences)
	}
}
