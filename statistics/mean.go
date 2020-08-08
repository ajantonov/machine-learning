package statistics

import "errors"

func Mean(values []float64) (float64, error) {
	size := len(values)
	isEmpty := 0 == size
	if isEmpty {
		return 0, errors.New("empty array")
	}

	var result = float64(0)

	for _, value := range values {
		result += value
	}
	result /= float64(size)

	return result, nil
}
