package statistics

import (
	"errors"
	"sort"
)

func TrimmedMean(values []float64, trimValue uint) (float64, error) {

	sizeOfValues := len(values)
	isInputArrayEmpty := 0 == sizeOfValues

	if isInputArrayEmpty {
		return 0, errors.New("invalid parameter values")
	}

	hasElementsIntoTrimmedArray := uint(sizeOfValues) <= uint(2*trimValue)
	if hasElementsIntoTrimmedArray {
		return 0, errors.New("trimValue is bigger than expected and trimmed array is empty")
	}

	var result = float64(0)

	sort.Float64s(values)

	trimmedValues := values[trimValue:(uint(sizeOfValues) - trimValue)]

	for _, value := range trimmedValues {
		result += value
	}

	sizeOfTrimmedValues := float64(len(trimmedValues))
	result /= sizeOfTrimmedValues

	return result, nil
}
