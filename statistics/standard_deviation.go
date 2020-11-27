package statistics

import "math"

/*
Standard deviation
	The square root of the variance.

	Synonyms
		l2-norm, Euclidean norm
*/
func StandardDeviation(samples []float64) (float64, error) {

	mean, meanError := Mean(samples)
	if meanError != nil {
		return 0, meanError
	}

	result, varianceError := Variance(samples, mean)
	if varianceError != nil {
		return 0, varianceError
	}

	result = math.Sqrt(result)

	return result, nil
}
