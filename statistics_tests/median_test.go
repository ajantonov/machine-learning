package statistics_tests

import (
	"github.com/ajantonov/machine-learning/statistics"
	"testing"
)

/**
Median Unit Tests

1. Median should return error when input array is empty
2. Median should return 1 when input array contains { 1 }
3. Median should return 6 when input array contains { 1, 3, 3, 6, 7, 8, 9 }
4. Median should return 6 when input array is contains unsorted elements { 1, 6, 3, 3, 7, 8, 9 }
5. Median should return 4.5 when input array contains {  1, 2, 3, 4, 5, 6, 8, 9 }
*/

func TestMedianShouldReturnErrorWhenInputArrayIsEmpty(t *testing.T) {

	var values []float64

	_, err := statistics.Median(values)

	if nil == err {
		t.Error("Median failed to return error when input array is empty!")
	}
}

func TestMedianShouldReturnTheOnlyElementIntoInputArray(t *testing.T) {

	var values = []float64{1}

	result, _ := statistics.Median(values)

	if !statistics.FloatEquals(1, result) {
		t.Errorf("Failed to calculate median for one element, expected 1 returned %f", result)
	}
}

func TestMedianShouldReturnMiddleValueElementWhenArrayIsOdd(t *testing.T) {

	values := []float64{1, 3, 3, 6, 7, 8, 9}

	result, _ := statistics.Median(values)

	if !statistics.FloatEquals(6, result) {
		t.Errorf("Median failed to calculate the median value expected 6 return %f", result)
	}
}

func TestMedianShouldReturnMiddleValueElementWhenArrayIsOddAndNotSorted(t *testing.T) {

	values := []float64{1, 6, 3, 3, 7, 8, 9}

	result, _ := statistics.Median(values)

	if !statistics.FloatEquals(6, result) {
		t.Errorf("Median failed to calculate the median value expected 6 return %f", result)
	}
}

func TestMedianShouldReturnMeanOfTheMiddleTwoElementWhenArrayIsEven(t *testing.T) {

	values := []float64{1, 2, 3, 4, 5, 6, 8, 9}

	result, _ := statistics.Median(values)

	if !statistics.FloatEquals(4.5, result) {
		t.Errorf("Median failed to calculate the median value expected 4.5 return %f", result)
	}
}
