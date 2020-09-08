package statistics

import "testing"

/*

Trimmed Mean Unit Tests

1. Weighted Mean should return error when input array is empty
2. Weighted Mean should return error when array with weight values is empty

*/
func TestWeightedMeanShouldReturnErrorWhenInputArrayIsEmpty(t *testing.T) {

	var inputs []float64
	var weights []float64

	_, err := WeightedMean(inputs, weights)

	if nil == err {
		t.Error("Failed to return error, when input array is empty!")
	}
}
