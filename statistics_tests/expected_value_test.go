package statistics_tests

import (
	"github.com/ajantonov/machine-learning/statistics"
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
		ExpectedValue Unit Tests
		Should return error when values array is empty - TestExpectedValueShouldReturnErrorValuesArrayIsEmpty
		Should return error when probabilities array is empty - TestExpectedValueShouldReturnErrorProbabilitiesArrayIsEmpty
		Should return error when values and probabilities arrays are with different length - TestExpectedValueShouldReturnErrorArraysAreWithDifferentLength
	    Should return valid result - TestExpectedValueShouldReturnValidResult
*/
func TestExpectedValueShouldReturnErrorValuesArrayIsEmpty(t *testing.T) {

	var values []float64
	var probabilities []float64 = []float64{1.0}

	_, err := statistics.ExpectedValue(values, probabilities)
	assert.Errorf(t, err, "ExpectedValue returned unexpected result")
}

func TestExpectedValueShouldReturnErrorProbabilitiesArrayIsEmpty(t *testing.T) {
	var values = []float64{1.0}
	var probabilities []float64

	_, err := statistics.ExpectedValue(values, probabilities)
	assert.Errorf(t, err, "ExpectedValue returned unexpected result")
}

func TestExpectedValueShouldReturnErrorArraysAreWithDifferentLength(t *testing.T) {
	var values = []float64{300, 50, 0}
	var probabilities = []float64{0.05, 0.15}

	_, err := statistics.ExpectedValue(values, probabilities)

	assert.EqualValuesf(t, statistics.ArraysAreWithDifferentLength, err, "ExpectedValue failed to return error ArraysAreWithDifferentLength")
}

func TestExpectedValueShouldReturnValidResult(t *testing.T) {
	var values = []float64{300, 50, 0}
	var probabilities []float64 = []float64{0.05, 0.15, 0.8}

	result, err := statistics.ExpectedValue(values, probabilities)

	assert.EqualValuesf(t, 22.5, result, "ExpectedValue failed to return valid result.")
	assert.NoError(t, err, "Returned Error is unexpected")
}
