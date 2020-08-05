package statistics

import "errors"

func Mean(values []float64) (float64, error) {
	if 0 == len(values) {
		return 0, errors.New("empty array")
	}

	var result = float64(0)
	var size = float64(len(values))

	for _, value := range values {
		result += value
	}
	result /= size

	return result, nil
}
