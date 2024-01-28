package statistics

import "fmt"

var (
	ValuesArrayIsEmptyError      = fmt.Errorf("values array is empty")
	ProbabilitiesArrayIsEmpty    = fmt.Errorf("probabilities array is empty ")
	ArraysAreWithDifferentLength = fmt.Errorf("both arrays are with different length")
)

func ExpectedValue(values []float64, probabilities []float64) (float64, error) {

	isValuesArrayEmpty := len(values) == 0
	if isValuesArrayEmpty {
		return 0, ValuesArrayIsEmptyError
	}

	isProbabilitiesArrayEmpty := len(probabilities) == 0
	if isProbabilitiesArrayEmpty {
		return 0, ProbabilitiesArrayIsEmpty
	}

	arraysAreWithDifferentLength := len(values) != len(probabilities)
	if arraysAreWithDifferentLength {
		return 0, ArraysAreWithDifferentLength
	}

	var result float64

	for index, value := range values {
		result += value * probabilities[index]
	}

	return result, nil
}
