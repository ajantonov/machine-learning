package statistics

import "testing"

/*

Trimmed Mean Unit Tests

1. Weighted Mean should return error when input array is empty
2. Weighted Mean should return error when array with weight values is empty
3. Weighted Mean should return error when sum of all weights is zero
4. Weighted Mean should return 1 when weights are { 1 } and input values { 1 }
5. Weighted Mean should return 2,666666667  when wights are { 1, 2, 3, 4, 5, 6 } and input values { 6, 5, 4, 3, 2, 1 }

*/

func TestWeightedMeanShouldReturnErrorWhenInputArrayIsEmpty(t *testing.T) {

	var inputs []float64
	var weights []float64

	_, err := WeightedMean(inputs, weights)

	if nil == err {
		t.Error("Failed to return error, when input array is empty!")
	}
}

func TestWeightedMeanShouldReturnErrorWhenWeightsArrayIsEmpty(t *testing.T) {

	var inputs = []float64{1}
	var weights []float64

	_, err := WeightedMean(inputs, weights)

	if nil == err {
		t.Error("Failed to return error, when weights array is empty!")
	}
}
