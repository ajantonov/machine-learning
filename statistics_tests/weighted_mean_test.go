package statistics_tests

import (
	"github.com/ajantonov/machine-learning/statistics"
	"testing"
)

/*

Trimmed Mean Unit Tests

1. Weighted Mean should return error when inputs array is empty
2. Weighted Mean should return error when array with weight values is empty
3. Weighted Mean should return error when sum of all weights is zero
4. Weighted Mean should return error when size of inputs and weights is different
5. Weighted Mean should return 1 when weights are { 1 } and input values { 1 }
6. Weighted Mean should return 2,666666667  when weights are { 1, 2, 3, 4, 5, 6 } and input values { 6, 5, 4, 3, 2, 1 }
*/

func TestWeightedMeanShouldReturnErrorWhenInputArrayIsEmpty(t *testing.T) {

	t.Parallel()

	var inputs []float64
	var weights []float64

	_, err := statistics.WeightedMean(inputs, weights)

	if nil == err {
		t.Error("Failed to return error, when input array is empty!")
	}
}

func TestWeightedMeanShouldReturnErrorWhenWeightsArrayIsEmpty(t *testing.T) {

	t.Parallel()

	var inputs = []float64{1}
	var weights []float64

	_, err := statistics.WeightedMean(inputs, weights)

	if nil == err {
		t.Error("Failed to return error, when weights array is empty!")
	}
}

func TestWeightedMeanShouldReturnErrorWhenSumOfAllWeightIsZero(t *testing.T) {

	t.Parallel()

	var inputs = []float64{1, 2, 3}
	var weights = []float64{1, -1, 0}

	_, err := statistics.WeightedMean(inputs, weights)

	if nil == err {
		t.Error("Failed to return error, when sum of all weights is empty!")
	}
}

func TestWeightedMeanShouldReturnErrorWhenSizeOfArraysIsDifferent(t *testing.T) {

	t.Parallel()

	var inputs = []float64{1, 2}
	var weights = []float64{1, 1, 0}

	_, err := statistics.WeightedMean(inputs, weights)

	if nil == err {
		t.Error("Failed to return error, when size of input arrays is different!")
	}
}

func TestWeightedMeanShouldReturnOne(t *testing.T) {

	t.Parallel()

	var inputs = []float64{1}
	var weights = []float64{1}

	result, _ := statistics.WeightedMean(inputs, weights)

	if !statistics.FloatEquals(1, result) {
		t.Errorf("WeightedMean expected to return 1, but returned %f ", result)
	}
}

func TestWeightedMeanShouldReturnTwoPointSix(t *testing.T) {

	t.Parallel()

	var inputs = []float64{6, 5, 4, 3, 2, 1}
	var weights = []float64{1, 2, 3, 4, 5, 6}

	result, _ := statistics.WeightedMean(inputs, weights)

	if !statistics.FloatEquals(2.666667, result) {
		t.Errorf("WeightedMean expected to return 2.666667 , but returned %f ", result)
	}
}
