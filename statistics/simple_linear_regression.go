package statistics

import "errors"

func LinearRegression(x []float64, y []float64) (slope float64, intercept float64, err error) {

	if 0 == len(x) {
		return float64(0), float64(0), errors.New("invalid parameter x")
	}

	if 0 == len(y) {
		return float64(0), float64(0), errors.New("invalid parameter y")
	}

	if len(x) != len(y) {
		return float64(0), float64(0), errors.New("different size of x and y")
	}

	// Calculate mean of x
	meanOfInputs, err := Mean(x)
	if nil != err {
		return float64(0), float64(0), errors.New("failed to calculate mean for x")
	}

	// Calculate mean of y
	meanOfOutputs, err := Mean(y)
	if nil != err {
		return float64(0), float64(0), errors.New("failed to calculate mean for y")
	}

	var numerator = float64(0)
	var denominator = float64(0)

	// Calculate slope
	for index := 0; index < len(x); index++ {
		numerator += (x[index] - meanOfInputs) * (y[index] - meanOfOutputs)
		denominator += (x[index] - meanOfInputs) * (x[index] - meanOfInputs)
	}

	slope = numerator / denominator

	// Calculate intercept
	intercept = meanOfOutputs - (slope * meanOfInputs)

	return slope, intercept, nil
}
