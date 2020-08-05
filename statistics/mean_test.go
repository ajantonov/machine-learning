package statistics

import (
	"testing"
)

/**
Mean Unit Tests
1. Should return error when array is empty  - TestMeanShouldReturnErrorWhenArrayIsEmpty
2. Should return result one when array has one element with value one - TestMeanShouldReturnOneWhenArrayContainsOne
3. Should return 2 when array has elements from 1 to 3 - func TestMeanShouldReturnTwoForRangeBetweenOneToThree
**/

func TestMeanShouldReturnErrorWhenArrayIsEmpty(t *testing.T) {
	var values = []float64{}

	_, err := Mean(values)

	if err == nil {
		t.Error("Mean function failed to return error, when array is empty! ")
	}
}

func TestMeanShouldReturnOneWhenArrayContainsOne(t *testing.T) {
	var values = []float64{1}

	result, _ := Mean(values)

	if float64(1) != result {
		t.Errorf("Mean function failed to calculate average of { 1 } = %f !", result)
	}
}

func TestMeanShouldReturnTwoForRangeBetweenOneToThree(t *testing.T) {
	var values = []float64{1, 2, 3}

	result, _ := Mean(values)

	if float64(2) != result {
		t.Errorf("Mean function failed to calculate average of { 1, 2, 3} = %f !", result)
	}
}
