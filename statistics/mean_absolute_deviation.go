package statistics

import (
	"errors"
	"math"
)

/*
Mean absolute deviation
	The mean of the absolute value of the deviations from the mean.
	Synonyms l1-norm, Manhattan norm
*/
func MeanAbsoluteDeviation(samples []float64, mean float64) (float64, error) {

	size := len(samples)
	if 0 == size {
		return 0, errors.New("error samples array is empty!")
	}

	var result = float64(0)

	for _, sample := range samples {
		result += math.Abs(sample - math.Abs(mean))
	}
	numberOfElements := float64(size)
	result /= numberOfElements

	return result, nil
}
