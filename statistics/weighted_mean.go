package statistics

import "errors"

func WeightedMean(inputs []float64, weights []float64) (float64, error) {
	if 0 == len(inputs) {
		return 0, errors.New("empty inputs parameter")
	}
	return 0, nil
}
