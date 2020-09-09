package statistics

import "errors"

/*
Weighted mean
	The sum of all values times a weight divided by the sum of the wieghts.

	Synonymous
		weighted average
*/
func WeightedMean(inputs []float64, weights []float64) (float64, error) {

	sizeOfInputs := len(inputs)
	if 0 == sizeOfInputs {
		return 0, errors.New("empty inputs parameter")
	}

	sizeOfWeights := len(weights)
	if 0 == sizeOfWeights {
		return 0, errors.New("empty weights parameter")
	}

	if sizeOfInputs != sizeOfWeights {
		return 0, errors.New(" two input arrays is with different length ")
	}

	var sumOfWeights = float64(0)
	for _, weight := range weights {
		sumOfWeights += weight
	}

	if float64(0) == sumOfWeights {
		return 0, errors.New("division by zero - sum of all weights is zero")
	}

	var sumOfEveryValueByWeight = float64(0)
	for index, value := range inputs {
		sumOfEveryValueByWeight += value * weights[index]
	}

	result := sumOfEveryValueByWeight / sumOfWeights

	return result, nil
}
