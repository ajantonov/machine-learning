package statistics

import "errors"

func WeightedMean(inputs []float64, weights []float64) (float64, error) {
	if 0 == len(inputs) {
		return 0, errors.New("empty inputs parameter")
	}

	if 0 == len(weights) {
		return 0, errors.New("empty weights parameter")
	}

	var sumOfWeights = float64(0)
	for _, weight := range weights {
		sumOfWeights += weight
	}

	if float64(0) == sumOfWeights {
		return 0, errors.New("division by zero - sum of all weights is zero")
	}

	return 0, nil
}
