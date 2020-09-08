package statistics

import (
	"testing"
)

/**
Trimmed Mean Unit Tests

*Agenda* p is smallest and largest values omitted

1. Trimmed Mean should return error when input array is empty.
2. Trimmed Mean should return error when input array contains { 0, 1 } and p = 1
3. Trimmed Mean should return error when input array contains { 0, 1, 2 } and p = 1
4. Trimmed Mean should return error when input array contains { 0, 1, 2 } and p = 2
5. Trimmed Mean should return 1 when input array contains { 0, 1, 2 } and p = 1
6. Trimmed Mean should return 2.5 when input array contains { 0, 2, 4, 3 } and p = 1
*/

func TestTrimmedMeanShouldReturnErrorWhenInputArrayIsEmpty(t *testing.T) {

	var values []float64
	trimValue := uint(0)

	_, err := TrimmedMean(values, trimValue)

	if err == nil {
		t.Error("TrimmedMean failed to return error, when input parameter is invalid!")
	}
}

func TestTrimmedMeanShouldReturnErrorWhenTrimValueIsOneAndSizeOfInputArrayIsTwo(t *testing.T) {

	var values = []float64{0, 1}
	trimValue := uint(1)

	_, err := TrimmedMean(values, trimValue)

	if nil == err {
		t.Error("TrimmedMean function should return error, when trim value is 1 and size of array is 2!")
	}
}

func TestTrimmedMeanShouldReturnOneWhenInputArrayContainsRangeOfZeroToTwoAndTrimValueIsTwo(t *testing.T) {

	var values = []float64{0, 1, 2}
	trimValue := uint(2)

	_, err := TrimmedMean(values, trimValue)

	if err == nil {
		t.Error("TrimmedMean function failed to return error, when trim value is too big regarding input array!")
	}
}

func TestTrimmedMeanShouldReturnOneWhenInputArrayContainsRangeOfZeroToTwoAndTrimValueIsOne(t *testing.T) {

	var values = []float64{0, 1, 2}
	trimValue := uint(1)

	result, err := TrimmedMean(values, trimValue)

	if float64(1) != result {
		t.Errorf("TrimmedMean function calculate %f for { 0, 1, 2 }, expected 1 ! Error %s ", result, err)
	}
}

func TestTrimmedMeanShouldReturnTwoPointFiveAndTrimValueIsOne(t *testing.T) {

	var values = []float64{0, 2, 4, 3}
	trimValue := uint(1)

	result, _ := TrimmedMean(values, trimValue)

	if 2.5 != result {
		t.Errorf("TrimmedMean function calculate %f for { 0, 2, 4, 3 }, expected 3", result)
	}

}
