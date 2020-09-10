package statistics

import (
	"errors"
	"sort"
)

/*
Median
	The value such that one-half of the data lies above and below

	Synonymous
		50th percentile
*/

func Median(values []float64) (float64, error) {

	sizeOfValues := len(values)

	if 0 == sizeOfValues {
		return 0, errors.New("invalid input parameter")
	}

	sort.Float64s(values)

	if 1 == sizeOfValues {
		return values[0], nil
	}

	middleIndex := sizeOfValues / 2

	hasArrayOddNumberOfElements := sizeOfValues%2 == 1
	if hasArrayOddNumberOfElements {
		return values[middleIndex], nil
	} else {
		meanOfMiddleTwoElements := (values[middleIndex-1] + values[middleIndex]) / 2
		return meanOfMiddleTwoElements, nil
	}
}
