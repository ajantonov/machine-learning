package statistics

import "errors"

/*
Variance
	The sum of differences between every element from samples and mean value divided by size of samples minus.
*/
func Variance(samples []float64, mean float64) (float64, error) {

	size := len(samples)
	if size <= 1 {
		return 0, errors.New("not enough elements")
	}

	var result float64

	for _, value := range samples {
		result += value - mean
	}
	result /= float64(size - 1)

	return result, nil
}
